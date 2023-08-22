package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email: "test_email@mail.com",
		Password: "password123",
	}
}