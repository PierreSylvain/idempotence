package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"idempotence/workflows"
	"idempotence/activities"
	
)


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading enrinoment variables file")
	}

	clientOptions := client.Options{
		HostPort: os.Getenv("TEMPORAL_HOST"),
		Namespace: os.Getenv("TEMPORAL_NAMESPACE"),
	}
    temporalClient, err := client.Dial(clientOptions)
    if err != nil {
        log.Fatal("Unable to create Temporal client", err)
    }
	defer temporalClient.Close()


    temporalWorker := worker.New(temporalClient, os.Getenv("TASKQUEUE"), worker.Options{})

	RegisterWFOptions := workflow.RegisterOptions{
		Name: "InventoryTask",
	}
	temporalWorker.RegisterWorkflowWithOptions(workflows.InventoryWorkflow, RegisterWFOptions)

	temporalWorker.RegisterActivity(activities.UpdateInventoryActivity)

    // Start listening to the task queue.
    err = temporalWorker.Run(worker.InterruptCh())
    if err != nil {
        log.Fatal("Unable to start worker", err)
    }
}
