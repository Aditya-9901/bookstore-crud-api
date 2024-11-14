package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:novell@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	gdb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	db = gdb
}

func GetDB() *gorm.DB {
	return db
}
