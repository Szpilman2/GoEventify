package routes

import (
	"goeventify/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message":"user created successfully"})
}

func Login(context *gin.Context)  {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message":"login successful"})
}