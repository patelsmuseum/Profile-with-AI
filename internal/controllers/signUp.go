package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SignUpController struct{}

func (ctrl SignUpController) SignUpPage(c *gin.Context) {
	c.HTML(200, "layout.html", gin.H{"Title": "Signup Page"})
}

func (ctrl SignUpController) SignUp(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println(email, password)
}
