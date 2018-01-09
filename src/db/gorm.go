package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlDb struct {
	User string
	Pass string
	Db   string
	Url  string
}

func init() {
	fmt.Println("initializing mysql db")
}

func (this *MysqlDb) Init() *gorm.DB {
	url := this.User + ":" + this.Pass + "@tcp(" + this.Url + ":3306)/" + this.Db + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(url)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	return db
}

func selectDBHost() string {
	if os.Getenv("DB_LINK") == "yes" {
		return "db"
	} else {
		return "localhost"
	}
}

func New() *gorm.DB {
	if os.Getenv("APP_ENV") == "development" {
		c := &MysqlDb{
			User: "root",
			Pass: "pass",
			Db:   "easycast",
			Url:  selectDBHost(),
		}
		db := c.Init()
		return db
	} else {
		c := &MysqlDb{
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Db:   os.Getenv("DB_NAME"),
			Url:  os.Getenv("DB_URL"),
		}
		db := c.Init()
		return db
	}
}
