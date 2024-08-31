package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muditsaxena1/url-shortner/internal/api"
)

func main() {
	port := flag.String("port", "8080", "server port")
	flag.Parse()

	router := gin.Default()

	api.SetupRoutes(router)

	log.Printf("Server running on port %s", *port)
	router.Run(":" + *port)
}
