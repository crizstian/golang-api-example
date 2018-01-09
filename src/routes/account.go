package routes

import "github.com/labstack/echo"

// internal function with company routes - demonstration only
func companyAPI(app *echo.Group) {
	// app.GET("/company/:account_id", api.GetUserCompany)
	// app.POST("/company/:account_id", api.SetUserCompany)
}

// Global function with local function and external functions, with api routes
func AccountAPI(app *echo.Group) {
	account := app.Group("/account")

	// external file example
	UserAPI(account)

	// internal function, with account group configuration
	companyAPI(account)
}
