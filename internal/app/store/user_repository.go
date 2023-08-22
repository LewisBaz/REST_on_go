package store

import "main/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encypted_password) VALUES ($1, $2) RETURNING id",
		user.Email,
		user.EncryptedPassword,
	).Scan(&user.ID); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encypted_password FROM users WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Email, &user.EncryptedPassword); err != nil {
		return nil, err
	}

	return user , nil
}