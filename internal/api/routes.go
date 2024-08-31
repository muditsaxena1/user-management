package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	// Basic Auth Middleware
	authorized := router.Group("/v1", gin.BasicAuth(gin.Accounts{
		"admin": "password", // Hardcoded credentials for testing
	}))

	authorized.GET("/user/:id", GetUser)
	authorized.POST("/user", SetUser)
	authorized.GET("/users", ListUsers)
}
