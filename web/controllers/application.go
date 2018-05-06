package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func notFound() mvc.View {
	return mvc.View{
		Code: iris.StatusNotFound,
		Name: "shared/404.html",
	}
}

func internalServerError() mvc.View {
	return mvc.View{
		Code: iris.StatusInternalServerError,
		Name: "shared/500.html",
	}
}
