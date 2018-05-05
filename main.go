package main

import (
	"github.com/adlerhsieh/gocasts/web/controllers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func irisApp() *iris.Application {
	app := iris.New()

	// Default config
	app.Use(recover.New())
	app.Use(logger.New())
	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.StaticWeb("/assets", "./public/assets")

	mvc.New(app).Handle(new(controllers.Screencasts))

	return app
}

func main() {
	app := irisApp()

	// Start the server using a network address.
	app.Run(iris.Addr(":8080"))
}
