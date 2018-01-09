package main

import (
	"easycast/src/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {

	// Setup
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/", nil)
	if assert.NoError(t, err) {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, api.IndexAPI(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, `{"Response":true}`, rec.Body.String())
		}
	}

}
