package main

import (
	"log"
	"os"
	"context"
	"go.temporal.io/sdk/client"
	"github.com/joho/godotenv"
	"idempotence/workflows"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading enrinoment variables file")
	}

	clientOptions := client.Options{
		HostPort: os.Getenv("TEMPORAL_HOST"),
		Namespace: os.Getenv("TEMPORAL_NAMESPACE"),
	}

	temporalClient, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer temporalClient.Close()
	
	workflowOptions := client.StartWorkflowOptions{
		ID: "inventory-task",
		TaskQueue:  os.Getenv("TASKQUEUE"),
	}

	workflowExec, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, workflows.InventoryWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", workflowExec.GetID(), "RunID", workflowExec.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = workflowExec.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}