package adapter

import "context"

/* __________________________________________________ */

type LoyaltyRequestModel struct {
	Order string
}

type LoyaltyResponseModel struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual,omitempty"`
}

/* __________________________________________________ */

type LoyaltyClient interface {
	Get(ctx context.Context, request LoyaltyRequestModel) (*LoyaltyResponseModel, error)
}

/* __________________________________________________ */
