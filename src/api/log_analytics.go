package api

import (
	ctrls "easycast/src/ctrls/analytics"
	"easycast/src/db"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var repo ctrls.LogAnalytics

// init functions are executed automatically one time, at startup time
// here we stablish a conection with our MongoDB cluster
func init() {
	session, err := db.MongoDB()
	if err != nil {
		SendError(ErrExt, ErrMsg["Mongo"], err)
	}
	db := session.DB("easycast")
	// making dependency injection
	repo = ctrls.InitLogAnalytics(db)
}

// getQueryParams is a function for the DRY pattern so we can abstract code
// and retrieve the necessary information from the headers of a request

func getQueryParams(c echo.Context, claim string) (string, string, string) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var entry string

	if claim == "account_entry" {
		entry = claims[claim].(string)
	} else {
		entry = strconv.FormatFloat(claims[claim].(float64), 'E', -1, 64)
	}

	startTime := c.QueryParam("startTime")
	endTime := c.QueryParam("endTime")

	return entry, startTime, endTime
}

// GetBandwidthByAccount ...
func GetBandwidthByAccount(c echo.Context) error {
	accountEntry, startTime, endTime := getQueryParams(c, "account_entry")
	result, err := repo.GetBandwidthByAccount(accountEntry, startTime, endTime)

	if err != nil {
		return SendError(ErrUsr, ErrMsg["Int"], err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"result":  result,
		"success": true,
	})
}

// GetLocationsByAccount ...
func GetLocationsByAccount(c echo.Context) error {
	accountEntry, startTime, endTime := getQueryParams(c, "account_entry")
	result, err := repo.GetLocationsByAccount(accountEntry, startTime, endTime)

	if err != nil {
		return SendError(ErrUsr, ErrMsg["Int"], err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"result":  result,
		"success": true,
	})
}
