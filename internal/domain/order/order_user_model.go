package order

import (
	"github.com/evgenivanovi/gomart/internal/domain/common"
	slices "github.com/evgenivanovi/gomart/pkg/std/slice"
)

/* __________________________________________________ */

type UserOrder struct {
	Owner common.UserID
	Order Order
}

func NewUserOrder(owner common.UserID, order Order) *UserOrder {
	return &UserOrder{
		Owner: owner,
		Order: order,
	}
}

/* __________________________________________________ */

type UserOrders struct {
	Owner  common.UserID
	Orders []Order
}

func NewUserOrders(owner common.UserID, orders []Order) *UserOrders {
	return &UserOrders{
		Owner:  owner,
		Orders: orders,
	}
}

func (o *UserOrders) HasOrders() bool {
	return slices.IsNotEmpty(o.Orders)
}

func (o *UserOrders) HasOrder(id OrderID) bool {
	filtered := slices.Filter(
		o.Orders,
		func(order Order) bool {
			return order.Identity().ID() == id.ID()
		},
	)
	return slices.IsNotEmpty(filtered)
}

func (o *UserOrders) ToUserOrders() []UserOrder {
	result := make([]UserOrder, 0)
	for _, order := range o.Orders {
		result = append(result, *NewUserOrder(o.Owner, order))
	}
	return result
}

/* __________________________________________________ */
