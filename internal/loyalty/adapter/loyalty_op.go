package adapter

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/order"
)

/* __________________________________________________ */

type LoyaltyOperationsService struct {
	client LoyaltyClient
}

func ProvideLoyaltyOperationsService(client LoyaltyClient) *LoyaltyOperationsService {
	return &LoyaltyOperationsService{
		client: client,
	}
}

func (l *LoyaltyOperationsService) Get(
	ctx context.Context, id order.OrderID,
) (*order.OrderLoyalty, error) {

	request := LoyaltyRequestModel{
		Order: id.ID(),
	}

	response, err := l.client.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	return l.toLoyalty(response)

}

func (l *LoyaltyOperationsService) toLoyalty(
	response *LoyaltyResponseModel,
) (*order.OrderLoyalty, error) {

	status, err := order.AsOrderStatus(response.Status)
	if err != nil {
		return nil, err
	}

	return order.NewOrderLoyalty(
		*order.NewOrderID(response.Order),
		*order.NewOrderData(status, response.Accrual),
	), nil

}

/* __________________________________________________ */
