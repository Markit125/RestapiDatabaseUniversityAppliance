package store

import "http-rest-api/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	AddStudentID(u *model.User, id int) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

type StudentRepository interface {
	Create(*model.Student) error
	Delete(*model.Student) error
	Find(int) (*model.Student, error)
	FindByPassport(string) (*model.Student, error)
}
