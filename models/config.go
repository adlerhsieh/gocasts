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
	// Why can't I use Where("key = ?", this.Key) ?
	db.DB.Where(&Config{Key: this.Key}).First(this)
}

func (this *Config) Save() {
	if this.ID == 0 {
		db.DB.Create(this)
	} else {
		db.DB.Save(this)
	}
}
