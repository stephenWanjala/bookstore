# Bookstore API
This repository contains a CRUD (Create, Read, Update, Delete) API for a bookstore implemented using the Go programming language (Golang) and the GORM ORM library with MySQL database.

The API provides endpoints for managing books, including creating new books, fetching book details, updating book information, deleting books. It follows RESTful principles and utilizes the Gorilla Mux router for handling HTTP requests.

## Requirements
- Go  1.21rc2 or higher
- MySQL 5.7 or higher
- Gorilla Mux 1.8.0 or higher
- GORM 1.9.16 or higher
- GoDotEnv 1.5.1 or higher
- Go MySQL Driver 1.5.0 or higher

## Installation
1. Clone the repository
2. Create a database in MySQL
3. Create a `.env` file in the root directory of the project and add the following environment variables:
    ```
    DB_USER=user
    DB_PASSWORD=secret
    DB_HOST=localhost
    DB_PORT=3306
    DB_NAME=bookstore
    ```

4. Run the application using `go run main.go`

API Endpoints
- GET `/api/books` - Get all books
- GET `/api/books/{id}` - Get a book by ID
- POST `/api/books` - Create a new book
- PUT `/api/books/{id}` - Update an existing book
- DELETE `/api/books/{id}` - Delete a book

## TODOs
- POST `/api/books/{id}/image` - Upload an image for a book 
-  Add pagination to GET `/api/books` endpoint
-  Add Relationships between models (e.g. Author, Publisher, etc.)

