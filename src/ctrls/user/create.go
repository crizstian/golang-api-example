package ctrls

import (
	"easycast/src/db"
	"easycast/src/errs"
	"easycast/src/libs"
	"easycast/src/models"
	"strconv"

	"github.com/labstack/echo"
)

func SetUserValues(c echo.Context) (interface{}, errs.Custom) {
	Db := db.New()
	defer Db.Close()

	id := c.Param("user_id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Bind"], err)
	}

	userProfile := new(models.UserProfile)

	err = c.Bind(userProfile)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Bind"], err)
	}

	user := new(models.User)
	err = user.FindByID(Db, ids)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["SQL"], err)
	}

	passChanged := false

	if userProfile.NewPass != "" {
		pass := libs.Password{}
		cp, err := pass.Compare(user.Pass, userProfile.CurrentPass)
		if err != nil {
			return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["Match"], err)
		}
		if cp {
			user.Pass, err = pass.Gen(userProfile.NewPass)
			if err != nil {
				return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["Password"], err)
			}
			Db.Save(&user)
			passChanged = true
		}
	}

	userInfo := new(models.UserInfo)

	userInfo.Name = userProfile.Name
	userInfo.LastName = userProfile.LastName
	userInfo.Phone = userProfile.Phone
	userInfo.UsersId = userProfile.UsersId

	ress := models.UserProfile{}
	ress.FindByID(Db, int(userProfile.UsersId))

	if ress.UsersId == 0 {
		Db.Save(&userInfo)
	} else {
		b := map[string]interface{}{
			"name":      userInfo.Name,
			"last_name": userInfo.LastName,
			"phone":     userInfo.Phone,
			"users_id":  userInfo.UsersId,
		}

		err := userInfo.Update(Db, int(userInfo.UsersId), b)
		if err != nil {
			return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["SQL"], err)
		}
	}

	if passChanged {
		res := map[string]interface{}{
			"msg":     "your data has been updated",
			"success": true,
		}
		return res, nil
	}

	res := map[string]interface{}{
		"msg":     "your data has been updated, Password not updated or modified",
		"success": true,
	}

	return res, nil
}
