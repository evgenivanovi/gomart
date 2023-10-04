package order

import "github.com/evgenivanovi/gomart/internal/model/order"

/* __________________________________________________ */

type OrderLoadRequestPayload struct {
	OrderID string
	UserID  int64
}

type OrderLoadRequest struct {
	Payload OrderLoadRequestPayload
}

/* __________________________________________________ */

type OrderLoadResponsePayload struct {
	Order order.UserOrder
}

type OrderLoadResponse struct {
	Payload OrderLoadResponsePayload
}

/* __________________________________________________ */
