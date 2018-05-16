package routes

import (
	"github.com/adlerhsieh/gocasts/web/handlers"

	"github.com/gin-gonic/gin"
)

func Draw(app *gin.Engine) {
	// Screencast
	app.GET("/", handlers.Screencast["index"])
	app.GET("/s/:slug", handlers.Screencast["show"])
	app.GET("/screencasts/new", handlers.Screencast["new"])
	app.POST("/screencasts", handlers.Screencast["create"])
	app.GET("/screencasts/edit/:slug", handlers.Screencast["edit"])
	app.POST("/screencasts/update/:slug", handlers.Screencast["update"])
	app.POST("/screencasts/destroy/:slug", handlers.Screencast["destroy"])

	// MISC
	app.GET("/author", handlers.Author["show"])

	// User
	app.GET("/signup", handlers.User["new"])
	app.POST("/users/create", handlers.User["create"])

	// Session
	app.GET("/signin", handlers.Session["new"])
	app.POST("/signin", handlers.Session["create"])
	app.POST("/signout", handlers.Session["destroy"])

	// Assets
	app.Static("/assets", "./web/assets")
}
