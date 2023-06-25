package router

import (
	"github.com/gorilla/mux"
	"github.com/stephenWanjala/bookstore/controllers"
	"gorm.io/gorm"
)

// SetupRouter sets up the router with routes and handlers
func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	bookController := &controllers.BookController{
		DB: db,
	}

	router.HandleFunc("/books", bookController.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", bookController.GetBookByID).Methods("GET")
	router.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")

	return router
}
