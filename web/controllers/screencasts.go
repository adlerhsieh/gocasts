package controllers

import (
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
	return mvc.View{
		Name: "screencasts/show.html",
		Data: iris.Map{
			"Layout": this.Layout,
			"id":     id,
		},
	}
}
