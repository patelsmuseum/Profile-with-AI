package main

import (
	"net/http"
	"os"
	"web-service-apis/internal/controllers"

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
	router.Static("/css", "internal/template/css")    //  /css will be served from internal/template/css
	router.Static("/jss", "internal/template/jss")    //  /jss will be served from internal/template/jss
	router.Static("/img", "internal/template/images") //  /img will be served from internal/template/img
	router.LoadHTMLGlob("internal/template/html/*")
	router.GET("/", func(c *gin.Context) {
		htmlContent, err := os.ReadFile("internal/template/html/homePage.html")
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
		signUpController := new(controllers.SignUpController)
		v1.GET("/welcome", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the API!"})
		})
		v1.GET("/signupPage", signUpController.SignUpPage)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/welcome", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the API2!"})
		})
	}

	router.Run(":8080")
}
