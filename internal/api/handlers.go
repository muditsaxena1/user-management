package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muditsaxena1/url-shortner/internal/models"
	"github.com/muditsaxena1/url-shortner/internal/storage"
)

var (
	UsersStorage = storage.NewInMemoryUserStorage()
)

func SetUser(c *gin.Context) {
	var user models.User

	// Bind the JSON body to the struct
	if err := c.ShouldBindJSON(&user); err != nil {
		errMessages := make(map[string]string)
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErrors {
				switch fieldErr.Field() {
				case "ID":
					if fieldErr.Tag() == "required" {
						errMessages["id"] = "id is required"
					}
				case "Name":
					if fieldErr.Tag() == "required" {
						errMessages["name"] = "Name is required"
					} else if fieldErr.Tag() == "min" {
						errMessages["name"] = "name must be at least 2 characters long"
					}
				case "SignupTime":
					if fieldErr.Tag() == "required" {
						errMessages["signupTime"] = "Signup time is required"
					} else if fieldErr.Tag() == "min" {
						errMessages["signupTime"] = "signupTime must be after 1 Jan 1850"
					}
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errMessages})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	if err := UsersStorage.SetUser(user); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user set successfully"})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	if user, err := UsersStorage.GetUser(id); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func ListUsers(c *gin.Context) {
	if users, err := UsersStorage.ListUsers(); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	} else {
		c.JSON(http.StatusOK, users)
	}
}
