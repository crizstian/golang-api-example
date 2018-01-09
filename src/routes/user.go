package routes

import (
	"easycast/src/api"

	"github.com/labstack/echo"
)

func UserAPI(app *echo.Group) {
	app.POST("/user", api.UserRegister)
	app.GET("/user", api.GetUsersList)
	app.PUT("/user", api.UpdateUserValues)
	app.DELETE("/user/:user_id", api.DeleteUser)

	app.GET("/user/profile/:user_id", api.GetUserData)
	app.PUT("/user/profile/:user_id", api.SetUserValues)
}
