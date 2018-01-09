package models

import (
	"easycast/src/db"
)

// AutoMigrate creates database tables
func AutoMigrate() {
	Db := db.New()
	defer Db.Close()

	Db.AutoMigrate(
		&Account{},
		&User{},
		&UserInfo{},
	)

}
