package main

import (
	"belajar-golang-jwt/src/helpers"
	"belajar-golang-jwt/src/entities"
	"belajar-golang-jwt/src/router"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if envLoadError := godotenv.Load(); envLoadError != nil {
		log.Fatal("[ERROR] Failed to load .env file")
	}
}

func main() {
	db := helpers.CreateDatabaseInstance()
	if migrateError := db.AutoMigrate(&entities.Food{}, &entities.User{}); migrateError != nil {
		log.Fatal("[ERROR] Couldn't migrate models!")
	}

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "9090"
	}

	router := router.RegisterRoutes(db)

	fmt.Printf("[OK] Server is started and listening on port: %s", port)

	log.Fatal(http.ListenAndServe(":" + port, router))
}
