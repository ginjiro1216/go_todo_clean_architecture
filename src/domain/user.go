package domain

import (
	"gorm.io/gorm"
)

type Users []User

type User struct {
	gorm.Model
	Todo     []Todo
	Password string
}
