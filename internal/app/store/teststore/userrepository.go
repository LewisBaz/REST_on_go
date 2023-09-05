package teststore

import (
	"main/internal/app/model"
	"main/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return  err
	}

	r.users[user.Email] = user
	user.ID = len(r.users)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	usr, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return usr, nil
}