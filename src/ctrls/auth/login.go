package ctrls

import (
	"easycast/src/db"
	"easycast/src/errs"
	"easycast/src/libs"
	"easycast/src/models"
	"errors"

	"github.com/labstack/echo"
)

func UserLogin(c echo.Context) (interface{}, errs.Custom) {
	// db connection
	Db := db.New()
	defer Db.Close()

	// user binding
	u := new(models.User)
	err := c.Bind(u)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Bind"], err)
	}
	// save pass plain first, u.Pass will change when we make "Db.First" with the password hashed from table user
	passPlain := u.Pass

	if u.Email == "" {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Login"], errors.New("Email not provided"))
	}

	if u.Pass == "" {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Login"], errors.New("Password not provided"))
	}

	// get the user by email
	count := 0
	err = u.FindAndCount(Db, &count)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrExt, errs.ErrMsg["SQL"], err)
	}
	if count == 0 {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Login"], errors.New("Email not match"))
	}
	if u.Active == 0 {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Account"], errors.New("Account inactive"))
	}

	// verify password
	pass := libs.Password{}
	cp, err := pass.Compare(u.Pass, passPlain)

	if err != nil {
		return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["Login"], err)
	}

	if !cp {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Login"], err)
	}

	// generate token
	j := new(libs.JWT)
	t, err := j.MapClaims(u.Email, u.ID)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["DecodingToken"], err)
	}

	res := map[string]interface{}{
		"token":   t,
		"msg":     "successfully logged in",
		"success": true,
	}

	return res, nil
}
