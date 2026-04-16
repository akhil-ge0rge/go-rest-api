package routes

import (
	"go-rest/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	//IF ONLY SINGLE ROUTE attach middle wares like this
	// server.POST("/events", middlewares.Authenticate, createEvents)

	//For multiple routes use grouped method

	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/events", createEvents)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.PATCH("/patch/events/:id", patchEvent)
	authenticate.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)

}
