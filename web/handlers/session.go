package handlers

import (
	"net/http"

	"github.com/adlerhsieh/gocasts/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var Session = map[string]func(c *gin.Context){
	"new": func(c *gin.Context) {
		c.HTML(http.StatusOK, "sessions/new", gin.H{})
	},
	"create": func(c *gin.Context) {
		email := c.PostForm("Email")
		user := models.User{}
		user.FindByEmail(email)

		password := c.PostForm("Password")
		err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password))
		if user.ID == 0 || err != nil {
			c.HTML(http.StatusOK, "sessions/new", gin.H{
				"msg": "Email or Password not matched",
			})
		} else {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	},
	"destroy": func(c *gin.Context) {

	},
}
