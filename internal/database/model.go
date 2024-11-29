package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `form:"name" binding:"required"`
	// email validate
	Email    string `gorm:"unique" binding:"required,email" form:"email"`
	Password string `validate:"binding" form:"password"`
	Avatar   string `form:"avatar"`
}
