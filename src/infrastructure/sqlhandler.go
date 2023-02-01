package infrastructure

import (
	"go_todo/interface/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	err      error
	dsn      = "host=db user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	dbConfig = &gorm.Config{}
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	db, err = gorm.Open(postgres.Open(dsn), dbConfig)
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.db.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.db.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.db.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.db.Save(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.db.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.db.Where(query, args...)
}
