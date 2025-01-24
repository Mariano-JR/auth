package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique;not null" json:"email" validate:"required,email"`
	Name     string `gorm:"not null" json:"name" validate:"required,min=3"`
	Password string `gorm:"not null" json:"password" validate:"required,min=6,max=18"`
}
