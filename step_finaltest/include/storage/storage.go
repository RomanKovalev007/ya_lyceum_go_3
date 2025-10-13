package storage

import (
	"final_project/include/models"
	"sync/atomic"
)

type Store struct{
	s map[int]models.User
	nextID int64
}

func New() *Store{
	return &Store{
		s: make(map[int]models.User),
		nextID: 1,
	}
}

func(s *Store) CreateUser(name string, age int) models.User {
	id := atomic.AddInt64(&s.nextID, 1) - 1
	createUser := models.User{ID: int(id), Name: name, Age: age}
	s.s[int(id)] = createUser
	return createUser
}

func(s *Store) GetUser(id int) (models.User, bool) {
	user, ok := s.s[id]
	return user, ok
}