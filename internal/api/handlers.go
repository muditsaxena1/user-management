package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muditsaxena1/url-shortner/internal/models"
	"github.com/muditsaxena1/url-shortner/internal/storage"
)

var (
	usersStorage = storage.NewInMemoryUserStorage()
)

func SetUser(c *gin.Context) {
	var user models.User

	// Bind the JSON body to the struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usersStorage.SetUser(user)

	c.JSON(http.StatusOK, gin.H{"status": "user set successfully"})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	if user, err := usersStorage.GetUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func ListUsers(c *gin.Context) {
	if users, err := usersStorage.ListUsers(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	} else {
		c.JSON(http.StatusOK, users)
	}
}
