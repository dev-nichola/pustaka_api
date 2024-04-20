package database

import (
	"pustaka_api/book"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() {
	dsn := "root:123@tcp(localhost:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&book.Book{})
}
