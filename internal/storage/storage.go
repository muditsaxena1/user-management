package storage

import (
	"github.com/muditsaxena1/url-shortner/internal/errors"
	"github.com/muditsaxena1/url-shortner/internal/models"
)

type Storage interface {
	SetUser(user models.User) *errors.CustomError
	GetUser(id string) (*models.User, *errors.CustomError)
	ListUsers() ([]models.User, *errors.CustomError)
}
