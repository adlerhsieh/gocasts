package main

import (
	"html/template"
	"time"

	"github.com/adlerhsieh/gocasts/db"
	"github.com/adlerhsieh/gocasts/models"
	"github.com/adlerhsieh/gocasts/web/routes"
	// "github.com/adlerhsieh/gocasts/web/controllers"

	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

// func irisApp() *iris.Application {
// 	app := iris.New()
//
// 	// Database Automigration
// 	db.DB.AutoMigrate(&models.Screencast{})
//
// 	// Config
// 	app.Use(recover.New())
// 	app.Use(logger.New())
//
// 	// View
// 	templates := iris.HTML("./web/views", ".html").Layout("shared/layout.html").Reload(true)
// 	templates.AddFunc("fmtDate", func(time time.Time) string {
// 		return time.Format("2006-01-02")
// 	})
// 	app.RegisterView(templates)
//
// 	// Assets
// 	app.StaticWeb("/assets", "./public/assets")
//
// 	// Routes & Controllers
// 	// https://godoc.org/github.com/kataras/iris/context#Context
// 	mvc.New(app.Party("/")).Handle(new(controllers.Screencast))
//
// 	return app
// }

func main() {
	app := gin.Default()

	// Database Automigration
	db.DB.AutoMigrate(&models.Screencast{})

	app.HTMLRender = viewConfig()

	routes.Draw(app)

	app.Run()
}

func viewConfig() *gintemplate.TemplateEngine {
	return gintemplate.New(gintemplate.TemplateConfig{
		Root:      "web/views",
		Extension: ".html",
		Master:    "shared/layout",
		Funcs: template.FuncMap{
			"fmtDate": func(time time.Time) string {
				return time.Format("2006-01-02")
			},
		},
		DisableCache: true,
	})

}
