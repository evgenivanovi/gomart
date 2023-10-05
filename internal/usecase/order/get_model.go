package order

import "github.com/evgenivanovi/gomart/internal/model/order"

/* __________________________________________________ */

type GetOrdersRequestPayload struct {
	UserID int64
}

type GetOrdersRequest struct {
	Payload GetOrdersRequestPayload
}

/* __________________________________________________ */

type GetOrdersResponsePayload struct {
	UserID int64
	Orders []order.Order
}

type GetOrdersResponse struct {
	Payload GetOrdersResponsePayload
}

/* __________________________________________________ */
