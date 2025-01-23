package user

import (
	"errors"

	"github.com/Mariano-JR/auth/internal/db"

	"github.com/google/uuid"
)

func Login(email, password string) (bool, error) {
	var user = db.DB.Where("email = ? AND password = ?", email, password).First(&User{})

	if user.Error == nil {
		return true, nil
	}

	return false, errors.New("invalid credentials")
}

func Save(email, name, password string) (bool, error) {
	user := User{
		ID:       uuid.NewString(),
		Email:    email,
		Name:     name,
		Password: password,
	}

	db.DB.Create(&user)

	return true, nil
}
