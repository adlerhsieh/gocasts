package main

import (
	"github.com/adlerhsieh/gocasts/db"
	"github.com/adlerhsieh/gocasts/models"
	"github.com/adlerhsieh/gocasts/web/controllers"

	// iris
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func irisApp() *iris.Application {
	app := iris.New()

	// Database Automigration
	db.DB.AutoMigrate(&models.Screencast{})

	// Config
	app.Use(recover.New())
	app.Use(logger.New())

	// View
	templates := iris.HTML("./web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(templates)

	// Assets
	app.StaticWeb("/assets", "./public/assets")

	// Routes & Controllers
	// https://godoc.org/github.com/kataras/iris/context#Context
	mvc.New(app.Party("/")).Handle(new(controllers.Screencast))

	return app
}

func main() {
	app := irisApp()

	app.Run(iris.Addr(":8080"))
}
