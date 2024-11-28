package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `form:"name" validate:"required"`
	// email validate
	Email    string `gorm:"unique" validate:"required,email" form:"email"`
	Password string `validate:"required" form:"password"`
	Avatar   string
}
