package handlers

import (
	"fmt"
	"net/http"

	"github.com/adlerhsieh/gocasts/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var Session = map[string]func(c *gin.Context){
	"new": func(c *gin.Context) {
		fmt.Println(c.MustGet("currentUser").(models.User))
		if c.MustGet("currentUser").(models.User).ID != 0 {
			c.Redirect(http.StatusFound, "/")
		}

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
			session := sessions.Default(c)
			session.Set("id", user.ID)
			session.Save()

			c.Redirect(http.StatusFound, "/")
		}
	},
	"destroy": func(c *gin.Context) {
		session := sessions.Default(c)

		fmt.Println("id", session.Get("id"))

		session.Delete("id")

		fmt.Println("id", session.Get("id"))

		err := session.Save()
		if err != nil {
			fmt.Println(err)
		}

		c.Redirect(http.StatusFound, "/")
	},
}
