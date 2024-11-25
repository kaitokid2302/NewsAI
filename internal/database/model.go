package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	// email validate
	Email    string `gorm:"unique" validate:"required,email"`
	Password string `validate:"required"`
	Avatar   string
}
