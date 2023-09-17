package teststore

import (
	"main/internal/app/model"
	"main/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return  err
	}

	user.ID = len(r.users) + 1
	r.users[user.ID] = user

	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	usr, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return usr, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}