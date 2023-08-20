package activities

import (
	"context"
	"go.temporal.io/sdk/activity"
)

func UpdateInventoryActivity(ctx context.Context) (error) {
	logger := activity.GetLogger(ctx)
	logger.Info("UpdateInventoryActivity")
	return nil
}
