package model_test

import (
	"main/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	test_cases := []struct{
		name string
		u func() *model.User
		isValid bool
	}{
		{
			name: "valid test",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				return &model.User{
					Email: "",
					Password: "password123",
				}
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				return &model.User{
					Email: "test_email@mail.com",
					Password: "",
				}
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				return &model.User{
					Email: "test_email@mail.com",
					Password: "1234",
				}
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				return &model.User{
					Email: "invalid",
					Password: "password123",
				}
			},
			isValid: false,
		},
		{
			name: "encrypt password",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""
				user.EncryptedPassword = "encryptedpassword"
				return user
			},
			isValid: true,
		},
	}

	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	user := model.TestUser(t)
	assert.NoError(t, user.BeforeCreate())
	assert.NotEmpty(t, user.EncryptedPassword)
}