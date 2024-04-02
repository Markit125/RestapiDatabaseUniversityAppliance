package teststore

import (
	"http-rest-api/internal/app/apiserver/store"
	"http-rest-api/internal/app/model"
)

// Store ...
type Store struct {
	userRepository    *UserRepository
	studentRepository *StudentRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[int]*model.User),
		}
	}

	return s.userRepository
}

// Student ...
func (s *Store) Student() store.StudentRepository {
	if s.studentRepository == nil {
		s.studentRepository = &StudentRepository{
			store:    s,
			students: make(map[int]*model.Student),
		}
	}

	return s.studentRepository
}
