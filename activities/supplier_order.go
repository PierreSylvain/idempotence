package activities

import (
	"context"
	"errors"
	"idempotence/inventory"
	"idempotence/tools"

	"go.temporal.io/sdk/activity"
)

// SupplierOrderActivity call supplier API to order new product.
//
// It takes a context.Context and an item string as parameters.
// It returns an error.
func SupplierOrderActivity(ctx context.Context, item string, quantity int) error {
	logger := activity.GetLogger(ctx)
	logger.Info("SupplierOrderActivity started")

	inStock, err := inventory.GetInStock(item)
	if err != nil {
		return err
	}
	if tools.IsError() {
		return errors.New("RANDOM ERROR")
	}
	if inStock < 10 {
		// Call supplier API and update inventory
		if tools.IsError() {
			return errors.New("RANDOM ERROR")
		}
		inventory.SupplierOrder(quantity)
		return nil
	}
	return nil
}
