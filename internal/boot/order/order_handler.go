package order

import (
	"github.com/evgenivanovi/gomart/internal/http/infra/order"
)

/* __________________________________________________ */

var (
	OrderLoadController *order.OrderLoadController
	GetOrdersController *order.GetOrdersController
)

/* __________________________________________________ */

func BootOrderHandlers() {
	OrderLoadController = order.ProvideOrderLoadController(
		OrderLoadUsecase,
	)
	GetOrdersController = order.ProvideGetOrdersController(
		GetOrdersUsecase,
	)
}

/* __________________________________________________ */
