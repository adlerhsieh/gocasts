package controllers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/adlerhsieh/gocasts/db"
	"github.com/adlerhsieh/gocasts/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type Screencast struct {
	Ctx    iris.Context
	Layout Layout
}

func (this *Screencast) BeginRequet(ctx iris.Context) {

}

func (this *Screencast) EndRequest(ctx iris.Context) {

}

func (this *Screencast) Get() mvc.View {
	screencasts := []models.Screencast{}
	db.DB.Find(&screencasts)

	fmt.Println(screencasts[0])

	return mvc.View{
		Name: "screencasts/index.html",
		Data: iris.Map{
			"Layout":      this.Layout,
			"screencasts": screencasts,
		},
	}
}

func (this *Screencast) GetBy(slug string) mvc.View {
	screencast := models.Screencast{}
	db.DB.Where("slug = ?", slug).First(&screencast)

	if screencast.ID == 0 {
		return notFound()
	}

	return mvc.View{
		Name: "screencasts/show.html",
		Data: iris.Map{
			"Layout":     this.Layout,
			"screencast": screencast,
		},
	}
}

func (this *Screencast) GetScreencastsNew() mvc.View {
	return mvc.View{
		Name: "screencasts/new.html",
		Data: iris.Map{
			"Layout": this.Layout,
		},
	}
}

func (this *Screencast) GetByEdit(slug string) mvc.View {
	screencast := models.Screencast{}
	db.DB.Where("slug = ?", slug).First(&screencast)

	if screencast.ID == 0 {
		return notFound()
	}

	return mvc.View{
		Name: "screencasts/edit.html",
		Data: iris.Map{
			"Layout":     this.Layout,
			"screencast": screencast,
		},
	}

}

func (this *Screencast) PostBy(slug string) mvc.Result {
	screencast := models.Screencast{}
	db.DB.Where("slug = ?", slug).First(&screencast)

	return mvc.Response{Path: "/" + screencast.Slug.String}
}

func bindFormToScreencast(ctx iris.Context) models.Screencast {
	screencast := models.Screencast{}
	if title := ctx.FormValue("Title"); title != "" {
		screencast.Title = sql.NullString{title, true}
	}
	if slug := ctx.FormValue("Slug"); slug != "" {
		screencast.Slug = sql.NullString{slug, true}
	}
	if abstract := ctx.FormValue("Abstract"); abstract != "" {
		screencast.Abstract = sql.NullString{abstract, true}
	}
	if content := ctx.FormValue("Content"); content != "" {
		screencast.Content = sql.NullString{content, true}
	}
	if video_embed := ctx.FormValue("VideoEmbed"); video_embed != "" {
		screencast.VideoEmbed = sql.NullString{video_embed, true}
	}
	if display_date := ctx.FormValue("DisplayDate"); display_date != "" {
		display_date, err := time.Parse("2006-01-02", display_date)
		if err != nil {
		} else {
			screencast.DisplayDate = display_date
		}
	}
	if thumbnail_url := ctx.FormValue("ThumbnailUrl"); thumbnail_url != "" {
		screencast.ThumbnailUrl = sql.NullString{thumbnail_url, true}
	}
	screencast.Public = (ctx.FormValue("Public") == "on")

	return screencast
}

func (this *Screencast) PostScreencasts() mvc.Result {
	screencast := bindFormToScreencast(this.Ctx)
	result := db.DB.Create(&screencast)

	if result.Error != nil {
		return mvc.View{
			Name: "screencasts/new.html",
			Data: iris.Map{
				"Layout": this.Layout,
				"alert":  result.Error,
			},
		}
	}

	return mvc.Response{Path: "/screencasts/" + screencast.Slug.String}
}
