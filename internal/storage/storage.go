package storage

import (
	"github.com/muditsaxena1/user-management/internal/errors"
	"github.com/muditsaxena1/user-management/internal/models"
)

type Storage interface {
	SetUser(user models.User) *errors.Error
	GetUser(id string) (*models.User, *errors.Error)
	ListUsers() ([]models.User, *errors.Error)
}
