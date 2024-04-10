package sqlstore

import (
	"database/sql"
	"errors"
	"http-rest-api/internal/app/apiserver/store"
	"http-rest-api/internal/app/model"
)

// StudentRepository ...
type StudentRepository struct {
	store *Store
}

// Create ...
func (r *StudentRepository) Create(s *model.Student) error {

	if err := s.Validate(); err != nil {
		return err
	}

	err := r.store.db.QueryRow(
		"INSERT INTO students "+
			"(first_name, middle_name, last_name, birth_date, achievements, passport) "+
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		s.FirstName,
		s.MiddleName,
		s.LastName,
		s.BirthDate,
		s.Achievements,
		s.Passport,
	).Scan(&s.ID)

	return err
}

// Delete ...
func (r *StudentRepository) Delete(s *model.Student) error {

	if err := s.Validate(); err != nil {
		return err
	}

	err := r.store.db.QueryRow(
		"DELETE FROM students WHERE id=$1",
		s.ID,
	).Err()

	return err
}

// FindByPassport ...
func (r *StudentRepository) FindByPassport(passport string) (*model.Student, error) {
	s := &model.Student{}
	if err := r.store.db.QueryRow(
		"SELECT id, passport FROM students WHERE passport=$1",
		passport,
	).Scan(
		&s.ID,
		&s.Passport,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return s, nil
}

// Find ...
func (r *StudentRepository) Find(id int) (*model.Student, error) {
	s := &model.Student{}
	if err := r.store.db.QueryRow(
		"SELECT id, first_name, middle_name, last_name, birth_date, achievements, passport FROM students WHERE id=$1",
		id,
	).Scan(
		&s.ID,
		&s.FirstName,
		&s.MiddleName,
		&s.LastName,
		&s.BirthDate,
		&s.Achievements,
		&s.Passport,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return s, nil
}
