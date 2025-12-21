package main

import (
	"fmt"
	"log"
	"rest-api/db"
	"rest-api/routes"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // loads .env from root
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.INITDB()

	fmt.Println("hello")
	r := gin.Default()
	fmt.Println(r)
	routes.RegisterEventsRoute(r)
	routes.RegisterUserRoutes(r)
	utils.GenerateJWT("a", 33)
	r.Run(":8080")
}
