package storage

import "github.com/muditsaxena1/url-shortner/internal/models"

type Storage interface {
	SetUser(user models.User) error
	GetUser(id string) (*models.User, error)
}
