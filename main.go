package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func SetupServer() *gin.Engine {
	r := gin.Default()
	r.GET("/", HomeHandler)
	return r
}

func main() {
	log.Println("Server Start...")
	SetupServer().Run()
}
