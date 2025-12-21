package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterEventsRoute(r *gin.Engine) {
	r.GET("/events", GetEvents)
	r.POST("/events", CreateEvent)
	r.GET("/event/:id", GetEventById)
	r.PUT("/event/:id", UpdateById)
	r.DELETE("/event/:id", DeleteById)

}
func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/signup", Signup)
	r.POST("/login", Login)
}
