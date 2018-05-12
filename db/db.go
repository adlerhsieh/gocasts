package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB = ConnectDB()

func ConnectDB() *gorm.DB {
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		database_url = "postgresql://hsiehadler@localhost:5432/gocasts?sslmode=disable"
	}

	// db, err := gorm.Open("mysql", "root:12345678@/gocasts?loc=Local&parseTime=true")
	db, err := gorm.Open("postgres", database_url)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	return db
}
