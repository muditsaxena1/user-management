package storage

import (
	"errors"
	"sync"

	"github.com/muditsaxena1/url-shortner/internal/models"
)

type InMemoryUserStorage struct {
	users map[string]models.User
	mutex sync.RWMutex
}

func NewInMemoryUserStorage() Storage {
	return &InMemoryUserStorage{users: make(map[string]models.User)}
}

func (s *InMemoryUserStorage) SetUser(user models.User) error {
	defer s.mutex.Unlock()
	s.mutex.Lock()
	s.users[user.ID] = user
	return nil
}

func (s *InMemoryUserStorage) GetUser(id string) (*models.User, error) {
	defer s.mutex.RUnlock()
	s.mutex.RLock()
	if user, found := s.users[id]; found {
		return &user, nil
	}
	return nil, errors.New("user not found")
}

func (s *InMemoryUserStorage) ListUsers() ([]models.User, error) {
	defer s.mutex.RUnlock()
	s.mutex.RLock()
	users := make([]models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users, nil
}
