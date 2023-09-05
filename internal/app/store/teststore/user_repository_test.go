package teststore_test

import (
	"main/internal/app/model"
	"main/internal/app/store"
	"main/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	str := teststore.New()
	usr := model.TestUser(t)

	assert.NoError(t, str.User().Create(usr))
	assert.NotNil(t, usr)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	str := teststore.New()

	email := "test_email@mail.com"
	_, err := str.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	usr := model.TestUser(t)
	usr.Email = email
	str.User().Create(usr)

	user, err := str.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
