package sqlstore_test

import (
	"http-rest-api/internal/app/apiserver/store"
	"http-rest-api/internal/app/apiserver/store/sqlstore"
	"http-rest-api/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudentRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("students")

	s := sqlstore.New(db)

	student := model.TestStudent(t)

	assert.NoError(t, s.Student().Create(student))
	assert.NotNil(t, student)
}

func TestStudentRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("students")

	s := sqlstore.New(db)

	student := model.TestStudent(t)

	s.Student().Create(student)

	err := s.Student().Delete(student)
	assert.NoError(t, err)
}

func TestStudentRepository_FindByPassport(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("students")

	s := sqlstore.New(db)

	passport := "1234567890"

	_, err := s.Student().FindByPassport(passport)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	student := model.TestStudent(t)
	s.Student().Create(student)

	u, err := s.Student().FindByPassport(student.Passport)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestStudentRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("students")

	s := sqlstore.New(db)

	_, err := s.Student().Find(-1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	student := model.TestStudent(t)
	s.Student().Create(student)

	u, err := s.Student().Find(student.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
