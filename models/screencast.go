package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Screencast struct {
	gorm.Model
	Slug         string         `gorm:"unique_index;not null"`
	Title        string         `gorm:"not null"`
	Content      sql.NullString `sql:"type:text"`
	Abstract     sql.NullString `sql:"type:text"`
	VideoEmbed   sql.NullString `sql:"type:text"`
	DisplayDate  time.Time      `sql:"type:date"`
	ThumbnailUrl sql.NullString
	Public       bool `gorm:"index:public"`
}
