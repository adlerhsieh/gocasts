package controllers

import (
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
		return mvc.View{
			Code: iris.StatusNotFound,
			Name: "shared/404.html",
		}
	}

	return mvc.View{
		Name: "screencasts/show.html",
		Data: iris.Map{
			"Layout": this.Layout,
			"title":  screencast.Title,
		},
	}
}
