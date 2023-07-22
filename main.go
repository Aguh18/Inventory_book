package main

import (

	"inventory/app"
	"inventory/db"

	"github.com/gin-gonic/gin"
)
var router *gin.Engine

func main() {
	conn := db.InitDB()
	router := gin.Default()
	router.LoadHTMLGlob("template/*")

	handler := app.New(conn)
	
	router.GET("books", handler.GetBooks)
	router.Run(":8080")

}

