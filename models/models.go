package models

import "gorm.io/gorm"

// Book represents a book in the bookstore
type Book struct {
	gorm.Model
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationDate string `json:"publication_date"`
}
