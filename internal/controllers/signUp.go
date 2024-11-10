package controllers

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type SignUpController struct{}

func (ctrl SignUpController) SignUpPage(c *gin.Context) {
	dir, err1 := filepath.Abs("./")
	if err1 != nil {
		fmt.Println(err1)
	}
	c.HTML(200, "signupPage.html", gin.H{"dir": dir})
}
