package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stephenWanjala/bookstore/models"
	"github.com/stephenWanjala/bookstore/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// BookController represents the controller for book-related operations
type BookController struct {
	DB *gorm.DB
}

// GetAllBooks returns all books
func (c *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	c.DB.Find(&books)
	utils.RespondJSON(w, http.StatusOK, books)
}

// GetBookByID returns a book by ID
func (c *BookController) GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var book models.Book
	c.DB.First(&book, id)

	if book.ID == 0 {
		utils.RespondError(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, book)
}

// CreateBook adds a new book
func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	c.DB.Create(&book)

	utils.RespondJSON(w, http.StatusCreated, book)
}

// UpdateBook updates an existing book
func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedBook models.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var book models.Book
	c.DB.First(&book, id)

	if book.ID == 0 {
		utils.RespondError(w, http.StatusNotFound, "Book not found")
		return
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author

	c.DB.Save(&book)

	utils.RespondJSON(w, http.StatusOK, book)
}

// DeleteBook deletes a book
func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var book models.Book
	c.DB.First(&book, id)

	if book.ID == 0 {
		utils.RespondError(w, http.StatusNotFound, "Book not found")
		return
	}

	c.DB.Delete(&book)

	var books []models.Book
	c.DB.Find(&books)

	utils.RespondJSON(w, http.StatusOK, books)
}
