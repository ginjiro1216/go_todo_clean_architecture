package domain

import (
	"gorm.io/gorm"
)

type Users []User

type User struct {
	gorm.Model
	Name     string `form:"name" json:"name"`
	Todo     []Todo
	Password string
}
