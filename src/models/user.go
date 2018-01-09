package models

import (
	"github.com/jinzhu/gorm"
)

type (
	// User represents the structure of our resource
	User struct {
		gorm.Model
		Email       string `json:"email" bson:"email" form:"email"`
		Pass        string `json:"pass" bson:"pass" form:"pass"`
		Role        int8   `json:"role" bson:"role" form:"role"`
		Active      int8   `json:"active" bson:"active" form:"active"`
		Token       string `json:"token" bson:"token" form:"token" sql:"type:text"`
		RecoverPass string `json:"recover" bson:"recover" form:"recover"`
		AccountId   uint   `json:"account_id" bson:"account_id" form:"account_id"`
	}

	// UserWithInfo used to send an email to account
	UserWithInfo struct {
		Email     string `json:"email"`
		Name      string `json:"name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		CreatedAt string `json:"created_at"`
	}
)

func (this *User) FindAll(Db *gorm.DB, accountID, offset, limit int) ([]User, error) {
	u := []User{}
	if err := Db.Offset(offset).Limit(limit).Where("account_id = ? AND active = 1", accountID).Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (this *User) FindByAccountID(Db *gorm.DB, accountId int) error {
	if err := Db.First(&this, "account_id = ?", accountId).Error; err != nil {
		return err
	}

	return nil
}

func (this *User) ExistEmail(Db *gorm.DB, email string) (bool, error) {
	// if err := Db.Where("email = ?", email).First(&this).Count(&count).Error; err != nil {
	// 	return true, err
	// }

	count := 0
	Db.Where("email = ?", email).First(&this).Count(&count)

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (this *User) FindByID(Db *gorm.DB, id int) error {
	if err := Db.Where("id = ?", id).First(&this).Error; err != nil {
		return err
	}

	return nil
}

// testing
// this.CreatedAt = time.Now()
// this.UpdatedAt = time.Now()
// this.ID = 4
// this.Email = "admin@easycast.com"
// this.Pass = "$2a$10$030owDRGDoJEutpxiriu3eypLXCheAyRObVai6hlA3HiP/.pB84aO"
// this.Active = 1
// this.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2VudHJ5IjoiIiwiYWNjb3VudF9pZCI6NCwiYWN0aXZlX3BsYW5faWQiOiIiLCJjcmVhdGVkSW50IjoyMDE3MDQsImN1c3RvbWVyX2lkIjoiY3VzXzJnVXpIenBLbURCNGJNbjc2IiwiZW1haWwiOiJhZG1pbkBlYXN5Y2FzdC5jb20iLCJleHAiOjE1MDE2NTcwNTcsImlhdCI6MTUwMTM5Nzg1NywicGFpZF9hdCI6IiIsInJvbGUiOjAsInVzZXJfaWQiOjR9.tiVkfVfmqwCAhENhG_3nQc7oJ5oaDDaJZpJqYt34IOk"
// this.AccountId = 4

// FindAndCount finds an user by email
func (this *User) FindAndCount(Db *gorm.DB, count *int) error {
	if result := Db.Where("email = ?", this.Email).First(&this).Count(count); result.Error != nil {
		if result.RecordNotFound() {
			return nil
		}
		return result.Error
	}

	return nil
}

func (this *User) UpdateToken(Db *gorm.DB, token string) error {
	this.Token = token
	if err := Db.Save(&this).Error; err != nil {
		return err
	}

	return nil
}

func (this *User) UpdateRecoverPass(Db *gorm.DB, pass, email string) error {
	if err := Db.Model(&this).Where("email = ?", email).Update("recover_pass", pass).Error; err != nil {
		return err
	}

	return nil
}

// GetTotalUser returns total records from asset based on account_id and asset_type
func (u *User) GetTotalUser(Db *gorm.DB, accountID int) int {
	count := 0
	Db.Select("id").Table("users").Where("account_id = ? AND active = 1", accountID).Count(&count)
	return count
}

// DeleteUser set active field to 0
func (u *User) DeleteUser(Db *gorm.DB, ID int) error {
	if result := Db.Table("users").Where("id = ?", ID).Update("active", 0); result.Error != nil {
		if result.RecordNotFound() {
			return result.Error
		}
	}
	return nil
}

// FindByCustomerID returns users info by customer_id (customers > users)
func (u *UserWithInfo) FindByCustomerID(Db *gorm.DB, customerID string) error {
	if result := Db.Select("users.email, users.created_at, users_info.name, users_info.last_name, users_info.phone").Table("customers").Joins("left join users on users.account_id = customers.account_id").Joins("left join users_info on users_info.users_id = users.id").Where("customers.customer_id = ? AND users.role = 2", customerID).First(&u); result.Error != nil {
		if result.RecordNotFound() {
			return nil
		}
		return result.Error
	}
	return nil
}
