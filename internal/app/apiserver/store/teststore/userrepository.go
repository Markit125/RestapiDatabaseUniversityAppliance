package teststore

import (
	"http-rest-api/internal/app/apiserver/store"
	"http-rest-api/internal/app/model"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
