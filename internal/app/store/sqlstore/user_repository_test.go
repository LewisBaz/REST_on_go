package sqlstore_test

import (
	"main/internal/app/model"
	"main/internal/app/store"
	"main/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	str := sqlstore.New(db)
	usr := model.TestUser(t)

	assert.NoError(t, str.User().Create(usr))
	assert.NotNil(t, usr)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	str := sqlstore.New(db)
	usr := model.TestUser(t)
	str.User().Create(usr)

	user, err := str.User().Find(usr.ID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	str := sqlstore.New(db)

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
