package controllers

import (
	"database/sql"
	// "fmt"

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
	return mvc.View{
		Name: "screencasts/index.html",
		Data: iris.Map{
			"Layout":  this.Layout,
			"message": "FUN",
		},
	}
}

func (this *Screencast) GetBy(id string) mvc.View {
	screencast := models.Screencast{}
	db.DB.Where("id = ?", id).First(&screencast)

	if screencast.ID == 0 {
		return notFound()
	}

	return mvc.View{
		Name: "screencasts/show.html",
		Data: iris.Map{
			"Layout": this.Layout,
			"title":  screencast.Title,
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

func (this *Screencast) GetByEdit(id string) mvc.View {
	screencast := models.Screencast{}
	db.DB.Where("id = ?", id).First(&screencast)

	if screencast.ID == 0 {
		return notFound()
	}

	return mvc.View{
		Name: "screencasts/edit.html",
		Data: iris.Map{
			"Layout": this.Layout,
			"title":  screencast.Title,
		},
	}

}

func bindFormToScreencast(ctx iris.Context) models.Screencast {
	screencast := models.Screencast{
		Title: ctx.FormValue("Title"),
		Slug:  ctx.FormValue("Slug"),
	}

	if abstract := ctx.FormValue("Abstract"); abstract != "" {
		screencast.Abstract = sql.NullString{abstract, true}
	}

	return screencast
}

func (this *Screencast) PostScreencasts() mvc.View {

	screencast := bindFormToScreencast(this.Ctx)

	// err := this.Ctx.ReadForm(&screencast)

	db.DB.Create(&screencast)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return internalServerError()
	// }

	return mvc.View{
		Name: "screencasts/index.html",
		Data: iris.Map{
			"Layout": this.Layout,
		},
	}
}
