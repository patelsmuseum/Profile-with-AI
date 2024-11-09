package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Hello, World!")
	defer logrus.Info("Goodbye, World!")
	router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{})
	// })
	router.GET("/", func(c *gin.Context) {
		htmlContent, err := os.ReadFile("internal/template/homePage.html")
		if err != nil {
			logrus.Error("Error reading home page template:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load home page"})
			return
		}
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, string(htmlContent))
	})
	v1 := router.Group("/v1")
	v1.Use()
	{
		v1.GET("/welcome", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the API!"})
		})
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/welcome", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the API2!"})
		})
	}

	router.Run(":8080")
}
