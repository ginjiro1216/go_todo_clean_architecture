package main

import (
	"go_todo/domain"
	"go_todo/infrastructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	err      error
	dsn      = "host=db user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	dbConfig = &gorm.Config{}
)

func main() {
	dbinit()
	_, err = gorm.Open(postgres.Open(dsn), dbConfig)
	if err != nil {
		panic(err)
	}
	infrastructure.Init()
}

func dbinit() {
	db, err = gorm.Open(postgres.Open(dsn), dbConfig)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(
		&domain.User{},
		&domain.Todo{},
	)
	if err != nil {
		panic(err)
	}
}
