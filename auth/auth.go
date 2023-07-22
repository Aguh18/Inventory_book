package auth

import (
	"inventory/models"
	"net/http"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context)  {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginHandler(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
	})
}
func LoginPostHandler(c *gin.Context)  {
	var  credetial models.Login

	err := c.Bind(&credetial)
	if  err !=nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "Username password invelid request",
		})
	}
	if credetial.Username != models.USER || credetial.Password != models.PASSWORD {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "Username password invelid request",
		}) 
	}
	else{
		claim := 
	}
	
}