package models

import (
	"fmt"

	"github.com/aditya-9901/bookstore/pkg/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b *Book) TableName() string {
	return "Book"
}

func init() {
	config.Connect()
	DB = config.GetDB()
	DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	DB.Create(b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	DB.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, error) {
	var book *Book
	err := DB.Where("ID=?", Id).Find(&book).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Book not found for the id")
		return nil, err
	}
	return book, nil
}

func DeleteBook(Id int64) *Book {
	var book *Book
	DB.Where("ID = ?", Id).First(&book)
	DB.Where("ID=?", Id).Delete(&book)
	return book
}
