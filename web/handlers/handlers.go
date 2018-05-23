package handlers

import (
	"github.com/adlerhsieh/gocasts/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func currentUser(c *gin.Context) models.User {
	session := sessions.Default(c)
	id := session.Get("id")

	if id != nil {
		user := models.User{}
		user.Find(int(id.(uint)))

		if user.Email.String != "" {
			return user
		}
	}

	return models.User{}
}
