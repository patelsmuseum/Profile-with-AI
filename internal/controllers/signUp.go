package controllers

import (
	"context"
	"fmt"
	"net/http"
	"web-service-apis/internal/database"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type SignUpController struct{}

func (ctrl SignUpController) SignUpPage(c *gin.Context) {
	c.HTML(200, "layout.html", gin.H{"Title": "Signup Page"})
}

func (ctrl SignUpController) SignUp(c *gin.Context) {
	c.Request.ParseForm()
	formData := c.Request.PostForm
	fmt.Println(formData)
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	firstName := c.PostForm("firstname")
	lastName := c.PostForm("lastname")

	confirmPassword := c.PostForm("confirmPassword")

	fmt.Println(email, password, phone, firstName, lastName, confirmPassword)

	mongoDb := database.GetMongoDB()
	userCollection := mongoDb.Database("Ai_profile").Collection("users") //user is collection Name
	var existingUser map[string]interface{}
	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&existingUser)
	if err != nil {
		log.Info("User not found Create new Account ", err)
	}
	fmt.Println("existingUser is ", existingUser)
	fmt.Println("monfodb connection is 3 ", mongoDb)

	newUser := bson.M{
		"email":     email,
		"firstName": firstName,
		"lastName":  lastName,
		"password":  password,
		"phone":     phone,
	}
	result, err := userCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		fmt.Println("Error in inserting user", err)
	}
	log.Info("User inserted successfully ", result.InsertedID)
	c.Redirect(http.StatusFound, "/v1/signinPage")

}
