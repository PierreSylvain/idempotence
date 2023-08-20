package main

import (
	"database/sql"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading enrinoment variables file")
	}

	// Remove the existing database file, if it exists
	os.Remove(os.Getenv("DATABASE"))

	// Open the database (will create if it doesn't exist)
	db, err := sql.Open("sqlite3", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the processedOrders table
	_, err = db.Exec(`CREATE TABLE processedOrders (
		orderID TEXT PRIMARY KEY
	)`)
	if err != nil {
		log.Fatalf("Failed to create processedOrders table: %s", err)
	}

	// Create the inventory table
	_, err = db.Exec(`CREATE TABLE inventory (
		productID TEXT PRIMARY KEY,
		quantity INTEGER
	)`)
	if err != nil {
		log.Fatalf("Failed to create inventory table: %s", err)
	}
}
