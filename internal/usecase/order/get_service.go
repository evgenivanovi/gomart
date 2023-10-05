package order

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/order"
	md "github.com/evgenivanovi/gomart/internal/model"
	order_md "github.com/evgenivanovi/gomart/internal/model/order"
)

/* __________________________________________________ */

type GetOrdersUsecase interface {
	Execute(context.Context, GetOrdersRequest) (GetOrdersResponse, error)
}

type GetOrderUsecaseService struct {
	transactor common.Transactor
	manager    order.OrderManager
}

func ProvideGetOrderUsecaseService(
	transactor common.Transactor,
	manager order.OrderManager,
) *GetOrderUsecaseService {
	return &GetOrderUsecaseService{
		transactor: transactor,
		manager:    manager,
	}
}

func (uc *GetOrderUsecaseService) Execute(
	ctx context.Context, request GetOrdersRequest,
) (GetOrdersResponse, error) {

	userID := common.NewUserID(request.Payload.UserID)

	ctx = uc.transactor.StartEx(ctx)
	ord, err := uc.manager.GetOrders(ctx, *userID)
	uc.transactor.CloseEx(ctx, err)

	if err != nil {
		return uc.toEmptyResponse(), err
	}

	if ord == nil {
		return uc.toEmptyResponse(), nil
	}

	return uc.toResponse(*ord), nil

}

func (uc *GetOrderUsecaseService) toEmptyResponse() GetOrdersResponse {
	return GetOrdersResponse{}
}

func (uc *GetOrderUsecaseService) toResponse(response order.UserOrders) GetOrdersResponse {

	orders := make([]order_md.Order, 0)

	for _, elem := range response.Orders {
		ord := order_md.Order{
			ID:      elem.Identity().ID(),
			Status:  elem.Data().Status.String(),
			Accrual: elem.Data().Accrual,
			Metadata: md.Metadata{
				CreatedAt: elem.Metadata().CreatedAt,
				UpdatedAt: elem.Metadata().UpdatedAt,
				DeletedAt: elem.Metadata().DeletedAt,
			},
		}
		orders = append(orders, ord)
	}

	return GetOrdersResponse{
		Payload: GetOrdersResponsePayload{
			UserID: response.Owner.ID(),
			Orders: orders,
		},
	}

}

/* __________________________________________________ */
