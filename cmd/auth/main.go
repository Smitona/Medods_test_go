package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Smitona/Medods_test_go/internal/models"
)

func main() {

	// connect to db
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// creating tables
	if err := db.AutoMigrate(&models.User{}, &models.Token{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// init routers
	router := AuthRouters()

	// init server
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
