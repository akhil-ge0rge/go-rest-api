package routes

import (
	"go-rest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userIdValue, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found in context",
		})
		return
	}

	userId, ok := userIdValue.(int64)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid userId type",
		})
		return
	}

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not able to parse the event id",
		})
		return
	}
	event, err := models.GetSingleEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not able to find the event",
		})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not able to register the event",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Registerd!",
	})

}

func cancelEventRegistration(context *gin.Context) {
	userIdValue, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found in context",
		})
		return
	}

	userId, ok := userIdValue.(int64)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid userId type",
		})
		return
	}

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not able to parse the event id",
		})
		return
	}

	var event models.Event

	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not able to cancel Registration",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Cancelled",
	})
}
