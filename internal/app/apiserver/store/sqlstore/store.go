package sqlstore

import (
	"database/sql"
	"http-rest-api/internal/app/apiserver/store"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db                *sql.DB
	userRepository    *UserRepository
	studentRepository *StudentRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}

// Student ...
func (s *Store) Student() store.StudentRepository {
	if s.studentRepository == nil {
		s.studentRepository = &StudentRepository{
			store: s,
		}
	}

	return s.studentRepository
}
