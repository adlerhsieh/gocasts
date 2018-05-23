package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type Screencast struct {
	gorm.Model
	Slug         sql.NullString `gorm:"unique_index;not null"`
	Title        sql.NullString `gorm:"not null"`
	Content      sql.NullString `sql:"type:text"`
	Abstract     sql.NullString `sql:"type:text"`
	VideoEmbed   sql.NullString `sql:"type:text"`
	DisplayDate  time.Time      `sql:"type:date"`
	ThumbnailUrl sql.NullString
	Public       bool `gorm:"index:public"`
}
