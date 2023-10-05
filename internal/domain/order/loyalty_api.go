package order

import "context"

/* __________________________________________________ */

type OrderLoyaltyOperations interface {
	Get(ctx context.Context, id OrderID) (*OrderLoyalty, error)
}

/* __________________________________________________ */
