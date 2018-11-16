package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupGin() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/health", func(c *gin.Context) {
		if health() {
			c.String(200, "OK")
		} else {
			c.String(500, "NOT OK")
		}
	})
	router.GET("/accounts/:id", func(c *gin.Context) {
		id := c.Param("id")
		accountImage, err := FindAccountImage(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, accountImage)
		}
	})
}
