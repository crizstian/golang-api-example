package db

import (
	"crypto/tls"
	"fmt"
	"net"

	mgo "gopkg.in/mgo.v2"
)

type MongoReplicaSet struct {
	User       string
	Pass       string
	Servers    string
	Db         string
	ReplicaSet string
	AuthSource string
}

var url MongoReplicaSet

func init() {
	fmt.Println("initializing mongodb db")
	url = MongoReplicaSet{
		User:       "cairo",
		Pass:       "cairo",
		Servers:    "cluster0-shard-00-00-a85hf.mongodb.net:27017,cluster0-shard-00-01-a85hf.mongodb.net:27017,cluster0-shard-00-02-a85hf.mongodb.net:27017",
		Db:         "easycast",
		ReplicaSet: "replicaSet=Cluster0-shard-0",
		AuthSource: "authSource=admin",
	}
}

func MongoDB() (*mgo.Session, error) {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
	}

	dialInfo, err := mgo.ParseURL("mongodb://" + url.User + ":" + url.Pass + "@" + url.Servers + "/" + url.Db + "?" + url.ReplicaSet + "&" + url.AuthSource)
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	return session, err
}
