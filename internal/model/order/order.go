package order

import md "github.com/evgenivanovi/gomart/internal/model"

/* __________________________________________________ */

type Order struct {
	ID       string
	Status   string
	Accrual  float64
	Metadata md.Metadata
}

type UserOrder struct {
	ID       string
	UserID   int64
	Status   string
	Accrual  float64
	Metadata md.Metadata
}

/* __________________________________________________ */
