package routes

import (
	"easycast/src/api"
	"easycast/src/libs"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// fns with first letter in uppercase, mean that are global functions inside
// the package

// ApiRoutes holds all the routes configurations, public as well as private
// all the routes must be grouped in a group route
func ApiRoutes(app *echo.Group) {
	// public routes - with external files
	AuthAPI(app)

	// restricted routes group
	restricted := makeRestrictedRoutes(app)

	// external configurations (global fn, that are in the same package)
	AccountAPI(restricted)
	AnalyticsAPI(restricted)
}

func makeRestrictedRoutes(app *echo.Group) *echo.Group {
	// restricted group
	j := new(libs.JWT)
	restricted := app.Group("/restricted")
	restricted.Use(middleware.JWT(j.Secret()))

	// inline routes - restricted test route
	restricted.GET("/test", api.IndexAPI)

	return restricted
}
