package main

import (
	"go_todo/domain"
	"go_todo/infrastructure/routers"
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
	dbInit()
	_, err = gorm.Open(postgres.Open(dsn), dbConfig)
	if err != nil {
		panic(err)
	}
	routers.Init()
}

func dbInit() {
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
