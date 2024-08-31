package storage

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/muditsaxena1/url-shortner/internal/errors"

	"github.com/muditsaxena1/url-shortner/internal/models"
)

type InMemoryUserStorage struct {
	users map[string]models.User
	mutex sync.RWMutex
}

func NewInMemoryUserStorage() Storage {
	return &InMemoryUserStorage{users: make(map[string]models.User)}
}

func (s *InMemoryUserStorage) SetUser(user models.User) *errors.CustomError {
	defer s.mutex.Unlock()
	s.mutex.Lock()
	s.users[user.ID] = user
	return nil
}

func (s *InMemoryUserStorage) GetUser(id string) (*models.User, *errors.CustomError) {
	defer s.mutex.RUnlock()
	s.mutex.RLock()
	if user, found := s.users[id]; found {
		return &user, nil
	}
	fmt.Printf("User:%s not found\n", id)
	return nil, errors.NewCustomError(http.StatusNotFound, "user not found")
}

func (s *InMemoryUserStorage) ListUsers() ([]models.User, *errors.CustomError) {
	defer s.mutex.RUnlock()
	s.mutex.RLock()
	users := make([]models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users, nil
}
