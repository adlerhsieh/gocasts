package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

func (this *Screencast) Path(action string) string {
	fmt.Println("123 " + action)
	switch action {
	case "index":
		return "/"
	case "show":
		return "/s/" + this.Slug.String
	case "new":
		return "/screencasts/new"
	case "create":
		return "/screencasts"
	case "edit":
		return "/screencasts/edit/" + this.Slug.String
	case "update":
		return "/screencasts/update/" + this.Slug.String
	case "destroy":
		return "/screencasts/destroy/" + this.Slug.String
	default:
		return ""
	}
}
