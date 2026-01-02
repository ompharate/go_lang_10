package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ompharate/go-bookstore-mysql/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:"name"`
	Author      string `gorm:"json:"author"`
	Publication string `gorm:"json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(book)
	return book
}

func (b *Book) UpdateBook(id int64) *Book {
	db.Model(&b).Where("ID = ?", id).Updates(Book{Name: b.Name, Author: b.Author, Publication: b.Publication})
	return b
}
