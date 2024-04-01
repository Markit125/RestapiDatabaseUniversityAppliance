package sqlstore_test

import (
	"http-rest-api/internal/app/apiserver/store"
	"http-rest-api/internal/app/apiserver/store/sqlstore"
	"http-rest-api/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(model.TestUser(t))

	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	_, err := s.User().Find(-1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u1 := model.TestUser(t)
	s.User().Create(u1)

	u, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
