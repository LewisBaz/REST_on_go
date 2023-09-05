package sqlstore

import (
	"database/sql"
	"main/internal/app/model"
	"main/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return  err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encypted_password) VALUES ($1, $2) RETURNING id",
		user.Email,
		user.EncryptedPassword,
	).Scan(&user.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encypted_password FROM users WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Email, &user.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		
		return nil, err
	}

	return user , nil
}