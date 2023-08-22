package main

import (
	"context"
	"fmt"
	"idempotence/inventory"
	"idempotence/workflows"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.temporal.io/sdk/client"
)

// main is the entry point of the program.
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading enrinoment variables file")
	}

	clientOptions := client.Options{
		HostPort:  os.Getenv("TEMPORAL_HOST"),
		Namespace: os.Getenv("TEMPORAL_NAMESPACE"),
	}

	temporalClient, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer temporalClient.Close()

	// Workflow options
	id := uuid.New()

	workflowOptions := client.StartWorkflowOptions{
		ID:        id.String(),
		TaskQueue: os.Getenv("TASKQUEUE"),
	}

	// Set values
	order := inventory.Order{
		OrderID:  "A12",
		Item:     "123456",
		Quantity: 999,
	}

	// Checking actual values
	inStock, err := inventory.GetInStock(order.Item)
	log.Printf("Product %s stock: %d\n", order.Item, inStock)

	// Execute workflow
	workflowExec, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, workflows.InventoryWorkflow, order)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", workflowExec.GetID(), "RunID", workflowExec.GetRunID())

	// Wait for the workflow completion.
	errWF := workflowExec.Get(context.Background(), nil)
	if errWF != nil {
		log.Fatalln("Unable get workflow result", errWF)
	}

	// Get product stock
	inStock, err = inventory.GetInStock(order.Item)
	fmt.Printf("Product %s stock: %d\n", order.Item, inStock)

}
