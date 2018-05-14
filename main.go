package main

import (
	"html/template"
	"time"

	"github.com/adlerhsieh/gocasts/db"
	"github.com/adlerhsieh/gocasts/models"
	"github.com/adlerhsieh/gocasts/web/routes"

	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
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
			// Global partials
			"shared/_navbar",

			// Form
			"screencasts/_form",

			// JS
			"screencasts/_highlight.js",

			// CSS
			"screencasts/_screencasts.css",
		},
		Funcs: template.FuncMap{
			"fmtDate": func(time time.Time) string {
				return time.Format("2006-01-02")
			},
			"toMarkdown": func(str string) string {

				// option := blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(
				// 	blackfriday.HTMLRendererParameters{
				// 		Flags: blackfriday.CommonHTMLFlags,
				// 	},
				// ))

				// return string(blackfriday.Run([]byte(str), option))
				return string(blackfriday.MarkdownCommon([]byte(str)))
			},
			"toHTML": func(str string) template.HTML {
				return template.HTML(str)
			},
		},
		DisableCache: true,
	})

}
