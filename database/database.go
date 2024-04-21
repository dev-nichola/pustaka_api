package database

import (
	"pustaka_api/book"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	dsn := "root:123@tcp(localhost:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&book.Book{})

	return db, err

	// Create Data

	// book := book.Book{
	// 	Title:       "Bumi Manusia",
	// 	Price:       90000,
	// 	Discount:    10,
	// 	Rating:      5,
	// 	Description: "Lorem ipsum dolor silit",
	// }

	// err = db.Create(&book).Error

	// if err != nil {

	// Mengambil Data
	// var books []book.Book

	// err = db.Where("title = ?", "Bumi Manusia").Find(&books).Error

	// if err != nil {
	// 	panic(err)
	// }

	// for _, book := range books {
	// 	fmt.Println("Title : ", book.Title)
	// 	fmt.Println("Book Object : %v", book)
	// }

	// Mengupdate Data
	// var book book.Book

	// err = db.Where("title = ?", "Bumi Manusia").Find(&book).Error

	// if err != nil {
	// 	panic(err)
	// }

	// book.Title = "Nichola Calon Programmer"

	// err = db.Save(&book).Error

	// if err != nil {
	// 	panic(err)
	// }

	// Menghapus Data
	// var book book.Book

	// err = db.Where("title = ?", "Bumi Manusia").Find(&book).Error

	// if err != nil {
	// 	panic(err)
	// }

	// err = db.Delete(&book, "id = 1").Error

	// if err != nil {
	// 	panic(err)
	// }

	// book := book.Book{
	// 	Title:       "Belajar Golang Dari Dasar Hingga Mahir",
	// 	Description: "Belajar Golang Pasti Kita Bisa",
	// 	Price:       9500000,
	// 	Rating:      5,
	// 	Discount:    10,
	// }

	// bookRepository.Create(book)

}
