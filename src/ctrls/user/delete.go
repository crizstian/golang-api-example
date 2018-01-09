package ctrls

import (
	"easycast/src/db"
	"easycast/src/errs"
	"easycast/src/models"
	"strconv"

	"github.com/labstack/echo"
)

// DeleteUser deletes an user
func DeleteUser(c echo.Context) (interface{}, errs.Custom) {
	Db := db.New()
	defer Db.Close()

	ID := c.Param("user_id")
	userID, _ := strconv.Atoi(ID)

	u := new(models.User)

	u.DeleteUser(Db, userID)

	res := map[string]interface{}{
		"id":      ID,
		"msg":     "New delete user",
		"success": true,
	}

	return res, nil
}
