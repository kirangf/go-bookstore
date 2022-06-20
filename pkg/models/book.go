package models

import (
	"github.com/jinzhu/gorm"
	appconfig "github.com/kirangf/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	appconfig.Connect()
	db = appconfig.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)

	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)

	return books
}

func GetBookById(bookID int) Book {
	var getBook Book
	db.Where("ID = ?", bookID).Find(&getBook)
	return getBook
}
