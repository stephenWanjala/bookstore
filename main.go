package main

import (
	"github.com/stephenWanjala/bookstore/models"
	"github.com/stephenWanjala/bookstore/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dsn := "root:wanjala@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal(err)
	}

	mRouter := router.SetupRouter(db)

	log.Fatal(http.ListenAndServe(":8000", mRouter))
}
