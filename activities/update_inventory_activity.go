package activities

import (
	"database/sql"
	"log"
	"os"
	"context"
	"go.temporal.io/sdk/activity"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func UpdateInventoryActivity(ctx context.Context) (error) {
	logger := activity.GetLogger(ctx)
	logger.Info("UpdateInventoryActivity started")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading enrinoment variables file")
	}

	db, err := sql.Open("sqlite3", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	orderID := "O12345"

	// Here it is the idempotence don't pprocess twice
	if orderIDExists(db, orderID) {
		return nil
	}
	addProcessedOrder(db, orderID)
	addInventoryItem(db, "P98765", 50)

	return nil
}

func orderIDExists(db *sql.DB, orderID string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(orderID) FROM processedOrders WHERE orderID = ?", orderID).Scan(&count)
	if err != nil {
		log.Fatalf("Failed to check orderID in processedOrders: %s", err)
	}
	return count > 0
}

func addProcessedOrder(db *sql.DB, orderID string) {
	_, err := db.Exec("INSERT INTO processedOrders (orderID) VALUES (?)", orderID)
	if err != nil {
		log.Fatalf("Failed to insert into processedOrders: %s", err)
	}
}

func addInventoryItem(db *sql.DB, productID string, quantity int) {
	_, err := db.Exec("INSERT INTO inventory (productID, quantity) VALUES (?, ?)", productID, quantity)
	if err != nil {
		log.Fatalf("Failed to insert into inventory: %s", err)
	}
}