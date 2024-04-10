package teststore

import (
	"http-rest-api/internal/app/apiserver/store"
	"http-rest-api/internal/app/model"
)

type StudentRepository struct {
	store    *Store
	students map[int]*model.Student
}

// Create ...
func (r *StudentRepository) Create(s *model.Student) error {
	if err := s.Validate(); err != nil {
		return err
	}

	s.ID = len(r.students) + 1
	r.students[s.ID] = s

	return nil
}

// Delete ...
func (r *StudentRepository) Delete(s *model.Student) error {
	delete(r.students, s.ID)
	return nil
}

// FindByEmail ...
func (r *StudentRepository) FindByPassport(passport string) (*model.Student, error) {
	for _, student := range r.students {
		if student.Passport == passport {
			return student, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// Find ...
func (r *StudentRepository) Find(id int) (*model.Student, error) {
	u, ok := r.students[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
