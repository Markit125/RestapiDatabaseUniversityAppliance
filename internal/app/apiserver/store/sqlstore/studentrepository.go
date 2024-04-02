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

	return r.store.db.QueryRow(
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
}

// FindByPassport ...
func (r *StudentRepository) FindByPassport(passport string) (*model.Student, error) {
	u := &model.Student{}
	if err := r.store.db.QueryRow(
		"SELECT id, passport FROM students WHERE passport=$1",
		passport,
	).Scan(
		&u.ID,
		&u.Passport,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

// Find ...
func (r *StudentRepository) Find(id int) (*model.Student, error) {
	u := &model.Student{}
	if err := r.store.db.QueryRow(
		"SELECT id, passport FROM students WHERE id=$1",
		id,
	).Scan(&u.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
