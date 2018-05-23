package models

import (
	"github.com/adlerhsieh/gocasts/db"

	"github.com/jinzhu/gorm"
)

type Config struct {
	gorm.Model
	Key   string
	Value string `sql:"type:text"`
}

func (this *Config) Get() {
	db.DB.Where("key = ?", this.Key).First(this)
}

func (this *Config) Save() {
	if this.ID == 0 {
		db.DB.Create(this)
	} else {
		db.DB.Save(this)
	}
}
