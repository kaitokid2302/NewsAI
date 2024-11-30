package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `form:"name" binding:"required" json:"name,omitempty"`
	// email validate
	Email    string `gorm:"unique" binding:"required,email" form:"email" json:"email,omitempty"`
	Password string `validate:"binding" form:"password" json:"password,omitempty"`
	Avatar   string `form:"avatar" json:"avatar,omitempty"`
}
