package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/len-mendonca/go-auth/models"
	"github.com/len-mendonca/go-auth/utils"
)

var validate *validator.Validate

func Login(c *gin.Context) {
	validate = validator.New()
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	//Check if email and password are filled
	if errorsMap := utils.ValidateEmailAndPassword(validate, &user); errorsMap != nil {
		c.JSON(400, gin.H{"error": errorsMap})
		return
	}

	var existingUser models.User

	if err := models.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	//  errHash := utils.CompareHashedPassword(user.Password,existingUser.Password);

	//  if(!errHash){
	// 	c.JSON(400,gin.H{"error":"Invalid password"})
	// 	return
	//  }

	//implementing jwt soon

}

func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	//var errHash error

	//user.Password,errHash = utils.GenerateHashPassword(user.Password)

	// if errHash != nil{
	// 	c.JSON(500,gin.H{"error":"Something went wrong in server"})
	// 	return
	// }

	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"success": "user created"})
	return

}

func Home(c *gin.Context) {

}

func Admin(c *gin.Context) {

}

func Logout(c *gin.Context) {

}
