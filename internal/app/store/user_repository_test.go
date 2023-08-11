package store_test

import (
	"main/internal/app/model"
	"main/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	str, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	user, err := str.User().Create(&model.User{
		Email: "test_email@mail.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	str, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email :=  "test_email@mail.com"
	_, err := str.User().FindByEmail(email)
	assert.Error(t, err)

	str.User().Create(&model.User{
		Email: "test_email@mail.com",
	})

	user, err := str.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}