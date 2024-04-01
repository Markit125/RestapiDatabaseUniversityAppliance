package teststore_test

import (
	"http-rest-api/internal/app/apiserver/store"
	"http-rest-api/internal/app/apiserver/store/teststore"
	"http-rest-api/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Create ...
func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

// FindByEmail ...
func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(model.TestUser(t))

	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// Find ...
func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()

	_, err := s.User().FindByEmail("unexisting@email.org")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	s.User().Create(u)

	u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
