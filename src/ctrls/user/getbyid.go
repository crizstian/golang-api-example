package ctrls

import (
	"easycast/src/db"
	"easycast/src/errs"
	"easycast/src/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetUserData(c echo.Context) (interface{}, errs.Custom) {
	Db := db.New()
	defer Db.Close()

	// here we are getting the id of the user passed in the route
	id := c.Param("user_id")

	ress := new(models.UserProfile)
	ids, _ := strconv.Atoi(id)
	err := ress.FindByID(Db, ids)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["Bind"], err)
	}

	res := map[string]interface{}{
		"user":    ress,
		"success": true,
	}

	return res, nil
}
