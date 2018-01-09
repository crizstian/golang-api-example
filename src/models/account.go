package models

import "github.com/jinzhu/gorm"

type (
	// User represents the structure of our resource
	Account struct {
		gorm.Model
		Name         string `json:"name"`
		Active       int    `json:"active"`
		LogoUrl      string `json:"logo_url"`
		Address      string `json:"address"`
		Phone        string `json:"phone"`
		VatNumber    string `json:"vatNumber"`
		CountriesId  uint   `json:"countries_id"`
		TimeZonesId  uint   `json:"time_zones_id"`
		Entry        string `json:"entry"`
		EntryTemp    string `json:"entry_temp"`
		ActivePlanId string `json:"active_plan_id"`
		Token        int    `json:"token"`
	}
)

// FindByID send account entry by id int
func (this *Account) FindByID(Db *gorm.DB, AccountId int) error {
	if err := Db.First(&this, "id = ?", AccountId).Error; err != nil {
		return err
	}

	return nil
}

// FindByIDAndCount send account entry by id int
func (this *Account) FindByIDAndCount(Db *gorm.DB, accountID int) (int, error) {
	count := 0

	if result := Db.Where("id = ?", accountID).Table("accounts").Count(&count); result.Error != nil {
		if result.RecordNotFound() {
			return 0, nil
		}
		return 0, result.Error
	}

	return count, nil
}

// GetAccountEntry send account entry by id int
func (this *Account) GetAccountEntry(Db *gorm.DB, accountID int) (string, error) {
	if result := Db.Where("ID = ?", accountID).First(&this); result.Error != nil {
		if result.RecordNotFound() {
			return "", nil
		}
		return "", result.Error
	}

	return this.Entry, nil
}

// FirstByUser ...
func (this *Account) FirstByUser(Db *gorm.DB, accountID uint) error {
	if err := Db.First(&this, "id = ?", accountID).Error; err != nil {
		return err
	}

	return nil
}
