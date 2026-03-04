package routes

import (
	"go-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message":     "Failed",
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Failed",
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message":     "User Created Sucessfuly",
		"status_code": http.StatusCreated,
		"data":        user,
	})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message":     "Failed",
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message":     "Failed",
			"status_code": http.StatusUnauthorized,
			"error":       err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message":     "Sucess",
		"status_code": http.StatusOK,
		"error":       nil,
	})
}
