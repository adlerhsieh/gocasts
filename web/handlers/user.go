package handlers

import (
	"database/sql"
	"net/http"

	"github.com/adlerhsieh/gocasts/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var User = map[string]func(c *gin.Context){
	"new": func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/new", gin.H{})
	},
	"create": func(c *gin.Context) {
		user := models.User{}
		msg := ""

		if email := c.PostForm("Email"); email != "" {
			user.Email = sql.NullString{email, true}
		} else {
			msg = "Email cannot be blank"
		}
		if username := c.PostForm("Username"); username != "" {
			user.Username = sql.NullString{username, true}
		} else {
			msg = "Username cannot be blank"
		}
		if password := c.PostForm("Password"); password != "" {
			pw, err := bcrypt.GenerateFromPassword([]byte(password), 10)
			if err != nil {
				msg = "Password generation error"
			} else {
				user.Password = sql.NullString{string(pw), true}
			}
		} else {
			msg = "Password cannot be blank"
		}

		if msg == "" {
			user.Save()
			c.Redirect(http.StatusMovedPermanently, "/signin")
		} else {
			c.HTML(http.StatusOK, "users/new", gin.H{
				"msg": msg,
			})
		}
	},
}
