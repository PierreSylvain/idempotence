package workflows

import (
	"time"
	"go.temporal.io/sdk/workflow"
	"idempotence/activities"
)
// InventoryWorkflow is a function that handles the inventory workflow.
// It takes a workflow context as input and returns an error if any.
func InventoryWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("InventoryWorkflow")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, activities.UpdateInventoryActivity).Get(ctx, nil)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return err
	}

	logger.Info("InventoryWorkflow completed.")

	return nil

}

