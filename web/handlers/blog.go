package handlers

import (
	"github.com/kataras/iris"
)

func Posts(ctx iris.Context) {
	ctx.View("posts/index.html")
}

func Post(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.ViewData("id", id)
	ctx.View("posts/show.html")
}
