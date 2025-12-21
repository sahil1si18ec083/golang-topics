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
	routes.RegisterEventsRoute(r)
	routes.RegisterUserRoutes(r)
	r.Run(":8080")
}
