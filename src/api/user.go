package api

import (
	ctrls "easycast/src/ctrls/user"
	"easycast/src/errs"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsersList(c echo.Context) error {
	res, err := ctrls.GetUsersList(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}

func GetUserData(c echo.Context) error {
	res, err := ctrls.GetUserData(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}

func SetUserValues(c echo.Context) error {
	res, err := ctrls.SetUserValues(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateUserValues(c echo.Context) error {
	res, err := ctrls.UpdateUserValues(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}

// DeleteUser endpoint that deletes a user
func DeleteUser(c echo.Context) error {
	res, err := ctrls.DeleteUser(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}
