package order

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/order"
	md "github.com/evgenivanovi/gomart/internal/model"
	order_md "github.com/evgenivanovi/gomart/internal/model/order"
)

/* __________________________________________________ */

type OrderLoadUsecase interface {
	Execute(context.Context, OrderLoadRequest) (OrderLoadResponse, error)
}

type OrderLoadUsecaseService struct {
	transactor common.Transactor
	manager    order.OrderManager
}

func ProvideOrderLoadUsecaseService(
	transactor common.Transactor,
	manager order.OrderManager,
) *OrderLoadUsecaseService {
	return &OrderLoadUsecaseService{
		transactor: transactor,
		manager:    manager,
	}
}

func (uc *OrderLoadUsecaseService) Execute(
	ctx context.Context, request OrderLoadRequest,
) (OrderLoadResponse, error) {

	userID := common.NewUserID(request.Payload.UserID)
	orderID := order.NewOrderID(request.Payload.OrderID)

	ctx = uc.transactor.StartEx(ctx)
	ord, err := uc.manager.LoadOrder(ctx, *userID, *orderID)
	uc.transactor.CloseEx(ctx, err)

	if err != nil {
		return uc.toEmptyResponse(), err
	}

	return uc.toResponse(*ord), err

}

func (uc *OrderLoadUsecaseService) toEmptyResponse() OrderLoadResponse {
	return OrderLoadResponse{}
}

func (uc *OrderLoadUsecaseService) toResponse(response order.UserOrder) OrderLoadResponse {

	return OrderLoadResponse{
		Payload: OrderLoadResponsePayload{
			Order: order_md.UserOrder{
				ID:      response.Order.Identity().ID(),
				UserID:  response.Owner.ID(),
				Status:  response.Order.Data().Status.String(),
				Accrual: response.Order.Data().Accrual,
				Metadata: md.Metadata{
					CreatedAt: response.Order.Metadata().CreatedAt,
					UpdatedAt: response.Order.Metadata().UpdatedAt,
					DeletedAt: response.Order.Metadata().DeletedAt,
				},
			},
		},
	}

}

/* __________________________________________________ */
