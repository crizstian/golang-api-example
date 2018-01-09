package libs

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/labstack/echo"
	// . "github.com/onsi/ginkgo"
	jwt "github.com/dgrijalva/jwt-go"
	. "github.com/onsi/gomega"
)

func ApiTest(api func(echo.Context) error, method string, url string, data interface{}, token interface{}, params map[string]interface{}) []byte {
	e := echo.New()
	var err error
	var req *http.Request

	if data != nil {
		req, err = http.NewRequest(method, url, strings.NewReader(data.(string)))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	Expect(err).ShouldNot(HaveOccurred())

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", token)

	if params != nil {
		u := params["url"].(string)
		values := params["values"].(map[string]string)

		c.SetPath(u)

		for key, _ := range values {
			c.SetParamNames(key)
			c.SetParamValues(values[key])
		}
	}

	// fapi := reflect.ValueOf(api)

	Expect(api(c)).ShouldNot(HaveOccurred())
	Expect(http.StatusOK).To(Equal(rec.Code))

	body, _ := ioutil.ReadAll(rec.Body)

	return body
}

func TokenTest(email string, role float64, account_id float64, id float64, customer_id float64) interface{} {
	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now()
	secs := now.Unix()

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = role
	claims["iat"] = secs
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["account_id"] = account_id
	claims["user_id"] = id
	claims["customer_id"] = customer_id

	return token
}

func TokenTestModify(account_entry string) interface{} {
	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now()
	secs := now.Unix()

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = "admin@easycast.com"
	claims["role"] = 1
	claims["iat"] = secs
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["account_id"] = 4.0
	claims["account_entry"] = account_entry

	return token
}
