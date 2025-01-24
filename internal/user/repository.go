package user

import (
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
	user, err := GetUser(email)

	if err == nil && user.Password == password {
		return true, nil
	}

	return false, err
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

func Delete(email string) (bool, error) {
	user, err := GetUser(email)

	if err != nil {
		return false, err
	}

	db.DB.Unscoped().Where("id = ?", user.ID).Delete(&User{})

	return true, nil
}
