package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	UserID uint64
	Note   string
}
