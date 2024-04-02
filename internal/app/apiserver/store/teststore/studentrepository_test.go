package teststore_test

import (
	"http-rest-api/internal/app/apiserver/store"
	"http-rest-api/internal/app/apiserver/store/teststore"
	"http-rest-api/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Create ...
func TestStudentRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestStudent(t)
	assert.NoError(t, s.Student().Create(u))
	assert.NotNil(t, u)
}

// FindByPassport ...
func TestStudentRepository_FindByPassport(t *testing.T) {
	s := teststore.New()

	passport := "1234567890"

	_, err := s.Student().FindByPassport(passport)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Student().Create(model.TestStudent(t))

	u, err := s.Student().FindByPassport(passport)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// Find ...
func TestStudentRepository_Find(t *testing.T) {
	s := teststore.New()

	_, err := s.Student().FindByPassport("unexisting_passport")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestStudent(t)
	s.Student().Create(u)

	u, err = s.Student().FindByPassport(u.Passport)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
