package activities

import (
	"context"
	"idempotence/inventory"

	"go.temporal.io/sdk/activity"
)

// CreateOrderActivity is a function that creates an order if the order does not exist.
//
// It takes a context.Context and an inventory.Order as parameters.
// It returns a boolean value true if the order was created other wise false and an error.
func CreateOrderActivity(ctx context.Context, order inventory.Order) (bool, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("CreateOrderActivity started")

	// Check if order exists
	orderExists := inventory.SearchOrder(order.OrderID)
	if orderExists {
		return false, nil
	}
	return true, nil
}
