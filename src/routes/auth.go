package routes

import (
	"easycast/src/api"

	"github.com/labstack/echo"
)

// Global function with authentication routes
func AuthAPI(app *echo.Group) {
	app.POST("/auth/login", api.UserLogin)
	app.POST("/auth/register", api.UserRegister)
}
