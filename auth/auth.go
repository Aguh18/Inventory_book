package auth

import (

	"inventory/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func HomeHandler(c *gin.Context)  {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginHandler(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content":"",
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
	} else{
		claim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute*5).Unix(),
			Issuer: "InventoryBook",
			IssuedAt: time.Now().Unix(),
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
		token, err := sign.SignedString([]byte(models.SECRET))
		if err != nil{
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"conttent": "username/password invalid",
			})
			c.Abort()
		}
		q := url.Values{}
		q.Set("auth", token)
		location := url.URL{Path: "/books", RawQuery: q.Encode()}
		c.Redirect(http.StatusMovedPermanently, location.RequestURI())
	}
	
	
}