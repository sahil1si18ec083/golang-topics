package main

import (
	"fmt"
	"rest-api/db"
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.INITDB()

	fmt.Println("hello")
	r := gin.Default()
	fmt.Println(r)
	r.GET("/events", routes.GetEvents)
	r.POST("/events", routes.CreateEvent)
	r.GET("/event/:id", routes.GetEventById)
	r.Run(":8080")
}
