package routes

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterEventsRoute(r *gin.Engine) {
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	protected.GET("/events", GetEvents)
	protected.POST("/events", CreateEvent)
	protected.GET("/event/:id", GetEventById)
	protected.PUT("/event/:id", UpdateById)
	protected.DELETE("/event/:id", DeleteById)

}

func RegisterUserRoutes(r *gin.Engine) {

	r.POST("/signup", Signup)
	r.POST("/login", Login)
}
