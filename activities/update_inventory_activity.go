package activities

import (
	"context"
	"idempotence/inventory"

	"go.temporal.io/sdk/activity"
)

// UpdateInventoryActivity updates the inventory based on the given order.
//
// Takes a context.Context and an inventory.Order as parameters.
// Returns an error.
func UpdateInventoryActivity(ctx context.Context, order inventory.Order) error {
	logger := activity.GetLogger(ctx)
	logger.Info("UpdateInventoryActivity started")

	// Check if order exists
	orderExists := inventory.SearchOrder(order.OrderID)
	if orderExists {
		return nil
	}
	inStock := inventory.GetInStock(order.Item)
	err := inventory.UpdateStock(order.OrderID, order.Item, inStock-order.Quantity)
	if err != nil {
		return err
	}
	return nil
}
