package main

import (
	"go-pra/copys/gin-copy/gin"
	"net/http"
)

func main() {
	e := gin.New()
	e.GET("/api/hello", Hello)
	e.Run("127.0.0.1:8080")
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hello": "World!",
	})
}
