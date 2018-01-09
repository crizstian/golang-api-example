package api

import (
	ctrls "easycast/src/ctrls/auth"
	"easycast/src/errs"
	"net/http"

	"github.com/labstack/echo"
)

// we follow the separation of concerns and MVC paradigm, in the api files, we only handle
// the request logic, the we forward it to the corresponding control to handle the logic or
// interactions with the db

func UserLogin(c echo.Context) error {
	res, err := ctrls.UserLogin(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}

func UserRegister(c echo.Context) error {
	res, err := ctrls.UserRegister(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}
