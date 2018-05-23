package routes

import (
	"github.com/adlerhsieh/gocasts/web/handlers"
	"github.com/adlerhsieh/gocasts/web/middlewares"

	"github.com/gin-gonic/gin"
)

func Draw(app *gin.Engine) {
	app.Use(middlewares.LoadCurrentUser)

	// Screencast
	app.GET("/", handlers.Screencast["index"])
	app.GET("/s/:slug", handlers.Screencast["show"])

	// MISC
	app.GET("/author", handlers.Author["show"])

	// User
	app.GET("/signup", handlers.User["new"])
	app.POST("/users/create", handlers.User["create"])

	// Session
	app.GET("/signin", handlers.Session["new"])
	app.POST("/signin", handlers.Session["create"])
	app.GET("/signout", handlers.Session["destroy"])

	// Admin
	admin := app.Group("/admin", middlewares.RequireAdmin)
	admin.GET("/screencasts/new", handlers.Screencast["new"])
	admin.POST("/screencasts", handlers.Screencast["create"])
	admin.GET("/screencasts/edit/:slug", handlers.Screencast["edit"])
	admin.POST("/screencasts/update/:slug", handlers.Screencast["update"])
	admin.POST("/screencasts/destroy/:slug", handlers.Screencast["destroy"])

	// Assets
	app.Static("/assets", "./web/assets")
}
