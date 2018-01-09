package ctrls

import (
	"easycast/src/db"
	"easycast/src/errs"
	"easycast/src/fn"
	"easycast/src/libs"
	"easycast/src/models"

	"github.com/labstack/echo"
)

// GetUsersList returns all users from an account
func GetUsersList(c echo.Context) (interface{}, errs.Custom) {
	Db := db.New()
	defer Db.Close()

	// get query params for page & limit
	offset, limit := libs.GetOffsetLimit(c.QueryParam("page"), c.QueryParam("limit"))

	// obtener el account_id desde el token
	accountID := fn.GetAccountID(c)

	u := new(models.User)
	userList, err := u.FindAll(Db, accountID, offset, limit)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["SQL"], err)
	}

	res := map[string]interface{}{
		"users":   userList,
		"success": true,
		"msg":     "list of users",
	}

	return res, nil
}
