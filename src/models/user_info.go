package models

import (
	"github.com/jinzhu/gorm"
)

type (
	UserInfo struct {
		gorm.Model
		Name     string `json:"name" bson:"name" form:"name"`
		LastName string `json:"lastname" bson:"lastname" form:"lastname"`
		Img      string `json:"img" bson:"img" form:"img"`
		Phone    string `json:"phone" bson:"phone" form:"phone"`
		UsersId  uint   `json:"users_id" bson:"users_id" form:"users_id"`
	}

	UserProfile struct {
		UsersId     uint   `json:"user_id"`
		Name        string `json:"name"`
		LastName    string `json:"lastname"`
		Phone       string `json:"phone"`
		CurrentPass string `json:"currentPass"`
		NewPass     string `json:"newPass"`
		NewPass2    string `json:"newPass2"`
	}

	UserUpdate struct {
		UsersId uint   `json:"users_id"`
		Pass    string `json:"pass"`
		Role    int8   `json:"role"`
		Active  int8   `json:"active"`
	}
)

func (UserInfo) TableName() string {
	return "users_info"
}

func (this *UserProfile) FindByID(Db *gorm.DB, id int) error {
	if result := Db.Select("name, last_name, phone, users_id").Table("users_info").Where("users_id = ?", id).Scan(&this); result.Error != nil {
		if result.RecordNotFound() {
			return nil
		}
		return result.Error
	}

	return nil
}

func (this *UserInfo) Update(Db *gorm.DB, id int, b map[string]interface{}) error {
	if err := Db.Model(&this).Where("users_id = ?", id).Updates(b).Error; err != nil {
		return err
	}

	return nil
}
