package handlers

import (
	"net/http"

	"github.com/adlerhsieh/gocasts/db"
	"github.com/adlerhsieh/gocasts/models"

	"github.com/gin-gonic/gin"
)

func ScreencastIndex(c *gin.Context) {
	screencasts := []models.Screencast{}
	db.DB.Find(&screencasts)

	c.HTML(http.StatusOK, "screencasts/index", gin.H{
		"screencasts": screencasts,
	})
}
