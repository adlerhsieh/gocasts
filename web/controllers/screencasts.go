package controllers

import (
	"github.com/kataras/iris/mvc"
)

type Screencasts struct {
}

func (c *Screencasts) GetFoo() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>FFF</h1>",
	}
}
