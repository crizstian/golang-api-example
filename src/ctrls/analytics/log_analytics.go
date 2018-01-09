package ctrls

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Example of dependecy injection in golang

// The Utils structure is used to initialized the db from the outside and
// be able to use one conecction obj, to make the corresponding requests,
// then the LogAnalytics interface is used to expose our internal functions,
// but when we use this interface from the outside we will have already the db
// obj created only one time and ready to execute the necessary logic

type (
	// Utils ...
	Utils struct {
		db *mgo.Database
	}

	// LogAnalytics ...
	LogAnalytics interface {
		GetBandwidthByAccount(string, string, string) ([]map[string]interface{}, error)
		GetLocationsByAccount(string, string, string) ([]interface{}, error)
	}
)

// InitLogAnalytics ...
func InitLogAnalytics(d *mgo.Database) LogAnalytics {
	u := new(Utils)
	u.db = d
	return u
}

// GetBandwidthByAccount ...
func (u Utils) GetBandwidthByAccount(customer string, startTime string, endTime string) ([]map[string]interface{}, error) {
	start, end, err := getTimeFormat(startTime, endTime)

	if err != nil {
		return nil, err
	}

	pipe := u.db.C("analytics").Pipe([]bson.M{
		{ // STAGE 1
			"$match": bson.M{
				"customer_id": customer,
				"timestamp": bson.M{
					"$elemMatch": bson.M{
						"$gte": start, "$lte": end,
					},
				},
			},
		}, { // STAGE 2
			"$group": bson.M{
				"_id":       bson.M{"customer_id": "$customer_id"},
				"bandwidth": bson.M{"$push": "$bytes"},
			},
		}, { // STAGE 3
			"$project": bson.M{
				"_id":          0,
				"accountEntry": "$_id.customer_id",
				"bandwidth":    bson.M{"$sum": "$bandwidth"},
			},
		},
	})

	return queryResponse(pipe)
}

// GetLocationsByAccount ...
func (u Utils) GetLocationsByAccount(customer string, startTime string, endTime string) ([]interface{}, error) {
	start, end, err := getTimeFormat(startTime, endTime)

	if err != nil {
		return nil, err
	}

	pipe := u.db.C("locations").Pipe(getBsonLocations(customer, start, end, "cip"))
	pipe2 := u.db.C("locations").Pipe(getBsonLocations(customer, start, end, "sip"))

	cip, err := queryResponse(pipe)
	if err != nil {
		return nil, err
	}
	sip, err2 := queryResponse(pipe2)
	if err2 != nil {
		return nil, err2
	}

	return []interface{}{
		cip,
		sip,
	}, nil
}

func getBsonLocations(customer string, start interface{}, end interface{}, loc string) []bson.M {
	return []bson.M{
		{
			"$match": bson.M{
				"customer_id": customer,
				"timestamp": bson.M{
					"$elemMatch": bson.M{
						"$gte": start, "$lte": end,
					},
				},
			},
		}, {
			"$group": bson.M{
				"_id": bson.M{"customer_id": "$customer_id"},
				loc:   bson.M{"$addToSet": "$" + loc},
			},
		}, {
			"$addFields": bson.M{
				loc: bson.M{
					"$reduce": bson.M{
						"input":        "$" + loc,
						"initialValue": make([]string, 0),
						"in":           bson.M{"$setUnion": []string{"$$value", "$$this"}},
					},
				},
			},
		}, {
			"$unwind": "$" + loc,
		}, {
			"$group": bson.M{
				"_id": bson.M{
					"customer_id": "$_id.customer_id",
					"country":     "$" + loc + ".country",
					"region":      "$" + loc + ".region",
					"city":        "$" + loc + ".city",
				},
				"total": bson.M{"$sum": 1},
				"info":  bson.M{"$addToSet": bson.M{"ll": "$" + loc + ".ll"}},
			},
		},
	}
}

func queryResponse(pipe *mgo.Pipe) ([]map[string]interface{}, error) {
	iter := pipe.Iter()

	var result []map[string]interface{}
	err := iter.All(&result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func getTimeFormat(startTime string, endTime string) (interface{}, interface{}, error) {
	var (
		start interface{}
		end   interface{}
		err   error
	)

	start, err = time.Parse(time.RFC3339, startTime)
	if err != nil {
		return "", "", err
	}
	end, err = time.Parse(time.RFC3339, endTime)
	if err != nil {
		return "", "", err
	}
	return start, end, err
}
