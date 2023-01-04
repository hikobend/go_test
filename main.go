package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" binding:"required"`
}

func router() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		// const StatusOK untyped int = 200
		// gin.Hはmap[string]interface{}と同じ
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})

	r.POST("/ps", func(c *gin.Context) {
		var u User
		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": u})
	})
	return r
}
func main() {
	router().Run()
}
