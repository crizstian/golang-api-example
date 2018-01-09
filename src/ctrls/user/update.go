package ctrls

import (
	"easycast/src/db"
	"easycast/src/errs"
	"easycast/src/libs"
	"easycast/src/models"

	"github.com/labstack/echo"
)

func UpdateUserValues(c echo.Context) (interface{}, errs.Custom) {
	Db := db.New()
	defer Db.Close()

	userUpdate := new(models.UserUpdate)

	err := c.Bind(userUpdate)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Bind"], err)
	}

	// modelo user
	u := new(models.User)

	// obtener informacion de "users" usando el UsersId

	err = u.FindByID(Db, int(userUpdate.UsersId))
	if err != nil {
		return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["SQL"], err)
	}

	u.Role, u.Active = userUpdate.Role, userUpdate.Active

	// si no son iguales los passwords (hash) el usuario lo actualizo
	if userUpdate.Pass != u.Pass {
		// hashear password
		pass := libs.Password{}
		u.Pass, err = pass.Gen(userUpdate.Pass)
		if err != nil {
			return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["Password"], err)
		}
	}

	// actualizar datos del usuario
	Db.Save(&u)

	res := map[string]interface{}{
		"msg": true,
	}

	return res, nil
}
