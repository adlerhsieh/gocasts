package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB = ConnectDB()

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:12345678@/gocasts?loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	return db
}
