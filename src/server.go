package main

import (
	"github.com/ginjiro1216/go_todo_clean_architecture/infrastructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=got_todo port=7777 sslmode=disable TimeZone=Asia/Tokyo"
	_, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	infrastructure.Init()
}
