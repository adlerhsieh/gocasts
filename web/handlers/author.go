package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Author = map[string]func(c *gin.Context){
	"show": func(c *gin.Context) {
		c.HTML(http.StatusOK, "author/show", gin.H{
			"currentUser": currentUser(c),
		})
	},
}
