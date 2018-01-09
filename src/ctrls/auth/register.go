package ctrls

import (
	"easycast/src/db"
	"easycast/src/errs"
	"easycast/src/libs"
	"easycast/src/models"

	"github.com/labstack/echo"
)

// struct to parse the json sent from the frontend
type UserRegisterStruct struct {
	Email     string `json:"email"`
	Pass      string `json:"pass"`
	Role      int8   `json:"role"`
	Active    int8   `json:"active"`
	AccountId uint   `json:"account_id"`
	Plan      string `json:"plan"`
}

// UserRegister registers a new user, returns an interface or an error
func UserRegister(c echo.Context) (interface{}, errs.Custom) {
	// db connection
	Db := db.New()
	// defer means that it will be the last line to be executed
	defer Db.Close()

	// create a UserRegisterStruct object to bind the json recieved
	ur := new(UserRegisterStruct)
	err := c.Bind(ur)
	if err != nil {
		// if an error occurred we send to the user the error,
		// errors could be from the frontend(user), internal(code), or external(db errors)
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Bind"], err)
	}

	// user model
	u := new(models.User)

	// check if email already exist, here the model handles the db interactions
	exist, err := u.ExistEmail(Db, ur.Email)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrExt, errs.ErrMsg["SQL"], err)
	}
	if exist {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["EmailTaken"], err)
	}

	// object assing to the user model
	u.Email = ur.Email
	u.Role = ur.Role
	u.Active = ur.Active

	// hashing password
	pass := libs.Password{}
	u.Pass, err = pass.Gen(ur.Pass)

	if err != nil {
		return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["CreateHash"], err)
	}

	// finally the user is inserted to the db
	Db.Create(&u)

	// finally we return a message to the frontend
	res := map[string]interface{}{
		"msg": "user registered",
	}

	return res, nil
}
