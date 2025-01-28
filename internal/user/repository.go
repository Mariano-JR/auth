package user

import (
	"errors"

	"github.com/Mariano-JR/auth/internal/db"

	"github.com/google/uuid"
)

func GetUser(email string) (*User, error) {
	var user User

	var err = db.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Login(email, password string) (bool, error) {
	if user, err := GetUser(email); err != nil {
		return false, err
	} else if user.Password == password {
		return true, nil
	}

	return false, errors.New("credentials invalids")
}

func Save(email, name, password string) (bool, error) {
	user := User{
		ID:       uuid.NewString(),
		Email:    email,
		Name:     name,
		Password: password,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func Delete(email string) (bool, error) {
	user, err := GetUser(email)

	if err != nil {
		return false, err
	}

	db.DB.Unscoped().Where("id = ?", user.ID).Delete(&User{})

	return true, nil
}
