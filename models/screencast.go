package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Screencast struct {
	gorm.Model
	Slug         string
	Title        string
	Content      string    `sql:"type:text"`
	Abstract     string    `sql:"type:text"`
	VideoEmbed   string    `sql:"type:text"`
	DisplayDate  time.Time `sql:"type:date"`
	ThumbnailUrl string
	Public       bool
}
