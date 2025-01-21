package user

import (
	"github.com/google/uuid"
)

var users = []User{}

func Save(Email, Name string) (bool, error) {
	user := User{
		ID:    uuid.NewString(),
		Email: Email,
		Name:  Name,
	}
	users = append(users, user)

	return true, nil
}

func Users() []User {
	return users
}
