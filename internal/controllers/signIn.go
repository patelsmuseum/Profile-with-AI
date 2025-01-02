package controllers

import (
	"fmt"
	"net/http"
	"web-service-apis/internal/database"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type SignInController struct{}

func (s *SignInController) SignInPage(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"Title": "Sign In Page",
	})
}

func (s *SignInController) SignIn(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	mongoDb := database.GetMongoDB()
	log.Info("MongoDb is connected to from signin controller", mongoDb)

	fmt.Println(email, password)
}
