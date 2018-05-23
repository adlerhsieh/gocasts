package main

import (
	"html/template"
	"time"

	"github.com/adlerhsieh/gocasts/db"
	"github.com/adlerhsieh/gocasts/models"
	"github.com/adlerhsieh/gocasts/web/routes"

	"github.com/foolin/gin-template"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("mysession", store))

	// Database Automigration
	db.DB.AutoMigrate(&models.Screencast{})
	db.DB.AutoMigrate(&models.User{})

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
			// Global partials
			"shared/_navbar",

			// Form
			"screencasts/_form",

			// JS

			// CSS
			"screencasts/_screencasts.css",
			"screencasts/_quill.css",
		},
		Funcs: template.FuncMap{
			"fmtDate": func(time time.Time) string {
				return time.Format("2006-01-02")
			},
			"toHTML": func(str string) template.HTML {
				return template.HTML(str)
			},
		},
		DisableCache: true,
	})

}
