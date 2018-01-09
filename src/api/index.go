package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func IndexAPI(c echo.Context) error {
	return c.Render(http.StatusOK, "doc", "")
}

func PingAPI(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func RobotsAPI(c echo.Context) error {
	return c.String(http.StatusOK, "User-agent: *\nDisallow: /")
}

func FaviconAPI(c echo.Context) error {
	return c.String(http.StatusNoContent, "")
}
