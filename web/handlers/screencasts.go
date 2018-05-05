package handlers

import (
	"github.com/kataras/iris"
)

func Screencasts(ctx iris.Context) {
	ctx.ViewData("message", "Hello")
	ctx.View("screencasts/index.html")
}

func Screencast(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.ViewData("id", id)
	ctx.View("screencasts/show.html")
}
