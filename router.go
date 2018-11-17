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
	router.GET("/accounts", func(c *gin.Context) {
		accountImages, err := ListAccountImages()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, accountImages)
		}
	})
	router.GET("/seed", func(c *gin.Context) {
		err := SeedRandomAccountImage()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"result": "ok"})
		}
	})
	router.Run(":8080")
}
