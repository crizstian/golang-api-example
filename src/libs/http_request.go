package libs

import (
	"bytes"
	"easycast/src/conf"
	"io/ioutil"
	"net/http"
	"strings"
)

func Http(source string, method string, url string, token string, data *bytes.Buffer) ([]byte, error) {
	var (
		u   string
		err error
		req *http.Request
	)

	switch source {
	case "Easypay":
		u = strings.Replace(url, source, conf.EasyPayHost, -1)
		break
	}

	if data != nil {
		req, err = http.NewRequest(method, u, data)
	} else {
		req, err = http.NewRequest(method, u, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if data != nil {
		req.Header.Set("Content-Length", "1")
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer"+" "+token)
	}

	return doReq(req)
}

func doReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
