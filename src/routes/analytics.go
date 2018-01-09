package routes

import (
	"easycast/src/api"

	"github.com/labstack/echo"
)

func AnalyticsAPI(app *echo.Group) {
	app.GET("/analytics/logs/bandwidth/", api.GetBandwidthByAccount)
	app.GET("/analytics/logs/locations/", api.GetLocationsByAccount)
}
