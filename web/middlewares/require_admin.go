package middlewares

import (
	"net/http"

	"github.com/adlerhsieh/gocasts/models"

	"github.com/gin-gonic/gin"
)

// This should be used after session, current_user
func RequireAdmin(c *gin.Context) {
	user := c.MustGet("currentUser").(models.User)

	if !user.IsAdmin() {
		c.Redirect(http.StatusFound, "/")
	}

	c.Next()
}
