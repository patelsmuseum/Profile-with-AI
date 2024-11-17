package main

import (
	"html/template"
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

	router.Static("/v1/css", "internal/template/css")    //  /css will be served from internal/template/css
	router.Static("/v1/jss", "internal/template/jss")    //  /jss will be served from internal/template/jss
	router.Static("/v1/img", "internal/template/images") //  /img will be served from internal/template/img
	router.SetHTMLTemplate(template.Must(template.New("layout").ParseFiles(
		"internal/template/html/layout.html",
		"internal/template/html/header.html",     // Base layout
		"internal/template/html/footer.html",     // Footer
		"internal/template/html/homePage.html",   // Home page content
		"internal/template/html/signupPage.html", // Signup page content
		"internal/template/html/signinPage.html", // Signin page content
	)))
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title": "Home Page",
		})
	})
	v1 := router.Group("/v1")
	v1.Use()
	{
		signUpController := new(controllers.SignUpController)
		v1.GET("/signupPage", signUpController.SignUpPage)
		v1.POST("/signup", signUpController.SignUp)
		signInController := new(controllers.SignInController)
		v1.GET("/signinPage", signInController.SignInPage)
		// v1.POST("/signin", signInController.SignIn)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/welcome", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the API2!"})
		})
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
