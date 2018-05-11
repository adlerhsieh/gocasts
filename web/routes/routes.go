package routes

import (
	"github.com/adlerhsieh/gocasts/web/handlers"

	"github.com/gin-gonic/gin"
)

func Draw(app *gin.Engine) {
	app.GET("/", handlers.Screencast["index"])
	app.GET("/:slug", handlers.Screencast["show"])
}
