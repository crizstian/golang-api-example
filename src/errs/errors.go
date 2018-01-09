package errs

import (
	"fmt"
	"net/http"

	"os"

	raven "github.com/getsentry/raven-go"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

var App_Env bool

func init() {
	if os.Getenv("APP_ENV") == "production" {
		App_Env = true
		raven.SetDSN(os.Getenv("SENTRY_DSN"))
	}
}

const (
	ErrUsr = "user"
	ErrExt = "external"
	ErrInt = "internal"
)

var ErrMsg = map[string]string{
	"SQL":                       "SQl syntax problem",
	"Query":                     "Query syntax",
	"Bind":                      "Can't bind data",
	"Write":                     "Can't write file",
	"LiveSelect":                "Can't select live channel",
	"AssetSelect":               "Can't select video entry",
	"IDne":                      "ID don't exist",
	"Upload":                    "Can't upload file",
	"BadInfo":                   "Some data is not provided",
	"CreateEntry":               "Can't create entry",
	"CreateCard":                "Can't create card",
	"DeleteCard":                "Can't delete card",
	"CreateHash":                "Can't create hash",
	"CreateCustomer":            "Can't create customer in the gateway payment",
	"FindCustomer":              "Can't find customer",
	"CreateInstance":            "Can't create instance",
	"UpdateInstance":            "Can't update instance",
	"CreateCharge":              "Can't create charge",
	"CycleCharge":               "Can't get cycle charge",
	"CreateSubscription":        "Can't create subscription",
	"UpdateSubscription":        "Can't update subscription",
	"PauseSubscription":         "Can't pause subscription",
	"CancelSubscription":        "Can't cancel subscription",
	"ResumeSubscription":        "Can't resume subscription",
	"CreateToken":               "Can't create token",
	"DecodingToken":             "Can't decode token",
	"UpdatePlan":                "Can't update account's plan",
	"Deletechannel":             "Can't delete instance",
	"SendEmail":                 "Can't send email",
	"Mongo":                     "Can't connect to mongodb",
	"UpdateAuthKey":             "Can't update auth key",
	"Login":                     "Email/password incorrect",
	"Account":                   "Account inactive",
	"EmailTaken":                "Email already taken",
	"Recover":                   "recover code does not exists",
	"CreateLive":                "Can't create live",
	"Match":                     "can't match password",
	"Password":                  "can't create password",
	"AccountFound":              "account not found",
	"DeleteLive":                "Can't delete live",
	"UpdateLive":                "Can't update live",
	"UpdateActivePlan":          "Can't update active plan in account",
	"RenameFolder":              "Can't rename folder",
	"RestoreFolder":             "Can't restore original folder name",
	"GetBandwidth":              "Can't get bandwidth",
	"GetStorage":                "Can't get storage",
	"GetMaximumStorage":         "Can't get maximum storage",
	"GetTrafficUsage":           "Can't GetTrafficUsage",
	"GetStreams":                "Can't get streams",
	"GetDataTransfered":         "Can't get data transfered",
	"GetDataTransferedInterval": "Can't get data transferred data with intervals",
	"IP2Location":               "Ip2Location error",
	"UnmarshalData":             "can't unmarshal data",
	"AccountAvailable":          "Account is already activated",
	"LimitsExceeded":            "Can't change plan, you are using more resources than allowed by the plan",
	"AccountExceeded":           "Can't get account limit exceeded",
}

func Send(status string, msg string, err error) *echo.HTTPError {
	m := msg
	var c int

	switch status {
	case "user":
		c = http.StatusBadRequest
		m += ", verify your data."
		log.Warn("An External error occured." + msg)
		break
	case "external":
		c = http.StatusInternalServerError
		m += " Something went wrong, please contact you're administrator."
		log.Warn("An External error occured." + msg)
		break
	case "internal":
		c = http.StatusNotAcceptable
		log.Warn("An Internal error occured." + msg)
		break
	}

	if App_Env == true {
		raven.CaptureError(err, nil)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("ERROR => ")
	log.Error(err)
	fmt.Println("-------------------------------------")
	fmt.Println("Sending Echo Error: " + m)
	fmt.Println("-------------------------------------")

	return echo.NewHTTPError(c, m)
}
