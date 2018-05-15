package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/adlerhsieh/gocasts/db"
	"github.com/adlerhsieh/gocasts/models"

	"github.com/gin-gonic/gin"
)

var Screencast = map[string]func(c *gin.Context){
	"index": func(c *gin.Context) {
		screencasts := []models.Screencast{}
		db.DB.Order("display_date desc").Find(&screencasts)

		c.HTML(http.StatusOK, "screencasts/index", gin.H{
			"screencasts": screencasts,
		})
	},
	"show": func(c *gin.Context) {
		screencast := models.Screencast{}
		db.DB.Where("slug = ?", c.Param("slug")).First(&screencast)

		if screencast.ID == 0 {
			c.Redirect(http.StatusMovedPermanently, screencast.Path("index"))
			return
		}

		c.HTML(http.StatusOK, "screencasts/show", gin.H{
			"screencast": screencast,
		})
	},
	"new": func(c *gin.Context) {
		c.HTML(http.StatusOK, "screencasts/new", gin.H{})
	},
	"create": func(c *gin.Context) {
		screencast := bindFormToScreencast(c, models.Screencast{})
		result := db.DB.Create(&screencast)

		if result.Error != nil {
			c.HTML(http.StatusOK, "screencasts/new", gin.H{
				"screencast": screencast,
			})
		}

		c.Redirect(http.StatusMovedPermanently, screencast.Path("show"))
	},
	"edit": func(c *gin.Context) {
		screencast := models.Screencast{}
		db.DB.Where("slug = ?", c.Param("slug")).First(&screencast)

		c.HTML(http.StatusOK, "screencasts/edit", gin.H{
			"screencast": screencast,
		})
	},
	"update": func(c *gin.Context) {
		screencast := models.Screencast{}
		db.DB.Where("slug = ?", c.Param("slug")).First(&screencast)

		screencast = bindFormToScreencast(c, screencast)
		result := db.DB.Save(&screencast)

		if result.Error != nil {
			c.HTML(http.StatusOK, "screencasts/new", gin.H{
				"screencast": screencast,
			})
		}

		c.Redirect(http.StatusMovedPermanently, screencast.Path("show"))
	},
	"destroy": func(c *gin.Context) {
		screencast := models.Screencast{}
		db.DB.Where("slug = ?", c.Param("slug")).First(&screencast)

		if screencast.ID == 0 {
		}
		db.DB.Delete(&screencast)

		c.Redirect(http.StatusMovedPermanently, screencast.Path("index"))
	},
}

func bindFormToScreencast(c *gin.Context, screencast models.Screencast) models.Screencast {
	if title := c.PostForm("Title"); title != "" {
		screencast.Title = sql.NullString{title, true}
	}
	if slug := c.PostForm("Slug"); slug != "" {
		screencast.Slug = sql.NullString{slug, true}
	}
	if abstract := c.PostForm("Abstract"); abstract != "" {
		screencast.Abstract = sql.NullString{abstract, true}
	}
	if content := c.PostForm("Content"); content != "" {
		screencast.Content = sql.NullString{content, true}
	}
	if video_embed := c.PostForm("VideoEmbed"); video_embed != "" {
		screencast.VideoEmbed = sql.NullString{video_embed, true}
	}
	if display_date := c.PostForm("DisplayDate"); display_date != "" {
		display_date, err := time.Parse("2006-01-02", display_date)
		if err != nil {
		} else {
			screencast.DisplayDate = display_date
		}
	}
	if thumbnail_url := c.PostForm("ThumbnailUrl"); thumbnail_url != "" {
		screencast.ThumbnailUrl = sql.NullString{thumbnail_url, true}
	}
	screencast.Public = (c.PostForm("Public") == "on")

	return screencast
}
