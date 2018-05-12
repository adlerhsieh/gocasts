package main

import (
	"html/template"
	"time"

	"github.com/adlerhsieh/gocasts/db"
	"github.com/adlerhsieh/gocasts/models"
	"github.com/adlerhsieh/gocasts/web/routes"

	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

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
		Partials: []string{
			"screencasts/_form",
		},
		Funcs: template.FuncMap{
			"fmtDate": func(time time.Time) string {
				return time.Format("2006-01-02")
			},
		},
		DisableCache: true,
	})

}
