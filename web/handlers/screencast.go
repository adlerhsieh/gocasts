package handlers

import (
	"net/http"

	"github.com/adlerhsieh/gocasts/db"
	"github.com/adlerhsieh/gocasts/models"

	"github.com/gin-gonic/gin"
)

var Screencast = map[string]func(c *gin.Context){
	"index": func(c *gin.Context) {
		screencasts := []models.Screencast{}
		db.DB.Find(&screencasts)

		c.HTML(http.StatusOK, "screencasts/index", gin.H{
			"screencasts": screencasts,
		})
	},
	"show": func(c *gin.Context) {
		screencast := models.Screencast{}
		db.DB.Where("slug = ?", c.Param("slug")).First(&screencast)

		c.HTML(http.StatusOK, "screencasts/show", gin.H{
			"screencast": screencast,
		})
	},
}
