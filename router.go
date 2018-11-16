package main

import "github.com/gin-gonic/gin"

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
	router.Run() // listen and serve on 0.0.0.0:8080
}
