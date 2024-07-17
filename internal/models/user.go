package models

import (
	"github.com/google/uuid"
)

// User represents a user entity.
type User struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// NewUser creates a new User instance.
func NewUser(username, password, firstName, lastName string) *User {
	return &User{
		UserID:    uuid.New().String(),
		Username:  username,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}
}
