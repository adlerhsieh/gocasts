package routes

import (
	"github.com/adlerhsieh/gocasts/web/handlers"

	"github.com/gin-gonic/gin"
)

func Draw(app *gin.Engine) {
	app.GET("/", handlers.Screencast["index"])
	app.GET("/s/:slug", handlers.Screencast["show"])
	app.GET("/screencasts/new", handlers.Screencast["new"])
	app.POST("/screencasts", handlers.Screencast["create"])
	app.GET("/screencasts/edit/:slug", handlers.Screencast["edit"])
	app.POST("/screencasts/update/:slug", handlers.Screencast["update"])
	app.POST("/screencasts/destroy/:slug", handlers.Screencast["destroy"])

	app.Static("/assets", "./web/assets")
}
