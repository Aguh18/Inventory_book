package main

import (
	"inventory/app"
	"inventory/auth"
	"inventory/db"

	"github.com/gin-gonic/gin"
)
var router *gin.Engine

func main() {
	conn := db.InitDB()
	router := gin.Default()
	router.LoadHTMLGlob("template/*")

	handler := app.New(conn)
	router.GET("/", auth.HomeHandler)

	router.GET("/login", auth.LoginHandler)

	router.POST("/home", auth.LoginPostHandler)

	router.GET("books", handler.GetBooks)
	router.Run(":8080")

}

