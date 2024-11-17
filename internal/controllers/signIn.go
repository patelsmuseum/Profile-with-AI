package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignInController struct{}

func (s *SignInController) SignInPage(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"Title": "Sign In Page",
	})
}
