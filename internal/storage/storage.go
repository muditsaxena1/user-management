package storage

import (
	"github.com/muditsaxena1/user-management/internal/errors"
	"github.com/muditsaxena1/user-management/internal/models"
)

type Storage interface {
	SetUser(user models.User) *errors.CustomError
	GetUser(id string) (*models.User, *errors.CustomError)
	ListUsers() ([]models.User, *errors.CustomError)
}
