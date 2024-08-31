package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muditsaxena1/url-shortner/internal/api"
)

func main() {
	router := gin.Default()

	api.SetupRoutes(router)

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
