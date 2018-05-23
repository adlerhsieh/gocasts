package middlewares

import (
	"github.com/adlerhsieh/gocasts/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoadCurrentUser(c *gin.Context) {
	c.Set("currentUser", currentUser(c))
	c.Next()
}

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
