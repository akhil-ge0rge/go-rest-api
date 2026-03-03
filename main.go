package main

import (
	"go-rest/db"
	"go-rest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)
	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Failed",
			"status_code": http.StatusInternalServerError,
			"error":       err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Sucess", "status_code": http.StatusOK, "data": events})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message":     "Failed",
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}
	event, err := models.GetSingleEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Failed",
			"status_code": http.StatusInternalServerError,
			"error":       err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message":     "Sucess",
		"status_code": http.StatusOK,
		"data":        event,
	})

}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message":     "Failed",
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}
	event.ID = 1
	event.UserID = 101
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Failed",
			"status_code": http.StatusInternalServerError,
			"error":       err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message":     "Sucess",
		"status_code": http.StatusCreated,
		"data":        event,
	})
}
