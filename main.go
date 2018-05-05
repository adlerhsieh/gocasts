package main

import (
	"github.com/adlerhsieh/gocasts/web/handlers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func irisApp() *iris.Application {
	app := iris.New()

	// Config
	app.Use(recover.New())
	app.Use(logger.New())
	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.StaticWeb("/assets", "./public/assets")

	// Screencasts
	app.Handle("GET", "/", handlers.Screencasts)
	app.Handle("GET", "/{id:string}", handlers.Screencast)

	// Blog posts
	app.Handle("GET", "/blog", handlers.Posts)
	app.Handle("GET", "/blog/{id:string}", handlers.Post)

	return app
}

func main() {
	app := irisApp()

	app.Run(iris.Addr(":8080"))
}
