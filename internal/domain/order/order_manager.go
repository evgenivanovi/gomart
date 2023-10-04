package order

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type OrderValidatorFn func(string) error

func (f OrderValidatorFn) And(other OrderValidatorFn) OrderValidatorFn {
	return func(value string) error {

		err := f(value)
		if err != nil {
			return err
		}

		err = other(value)
		if err != nil {
			return err
		}

		return nil

	}
}

type OrderValidator interface {
	Validate(string) error
}

type OrderValidatorService struct {
	validators []func(string) error
}

func ProvideOrderValidatorService(
	validators ...func(string) error,
) *OrderValidatorService {
	return &OrderValidatorService{
		validators: validators,
	}
}

func (o OrderValidatorService) Validate(value string) error {
	for _, validator := range o.validators {
		err := validator(value)
		if err != nil {
			return errx.NewErrorWithEntityCode(ErrorOrderEntity, ErrorOrderInvalid)
		}
	}
	return nil
}

/* __________________________________________________ */

type OrderManager interface {
	LoadOrder(
		ctx context.Context, owner common.UserID, order OrderID,
	) (*UserOrder, error)

	GetOrders(
		ctx context.Context, owner common.UserID,
	) (*UserOrders, error)
}

type OrderManagerService struct {
	loyaltyOperations    OrderLoyaltyOperations
	orderValidator       OrderValidator
	userOrderRepository  UserOrderRepository
	userOrdersRepository UserOrdersRepository
}

func ProvideOrderManagerService(
	loyaltyOperations OrderLoyaltyOperations,
	orderValidator OrderValidator,
	userOrderRepository UserOrderRepository,
	userOrdersRepository UserOrdersRepository,
) *OrderManagerService {
	return &OrderManagerService{
		loyaltyOperations:    loyaltyOperations,
		orderValidator:       orderValidator,
		userOrderRepository:  userOrderRepository,
		userOrdersRepository: userOrdersRepository,
	}
}

func (o *OrderManagerService) LoadOrder(
	ctx context.Context, owner common.UserID, order OrderID,
) (*UserOrder, error) {

	err := o.orderValidator.Validate(order.ID())
	if err != nil {
		return nil, err
	}

	err = o.checkOrderExistence(ctx, owner, order)
	if err != nil {
		return nil, err
	}

	ord, err := o.associateOrder(ctx, owner, order)
	if err != nil {
		return nil, err
	}

	return ord, nil

}

func (o *OrderManagerService) GetOrders(
	ctx context.Context, owner common.UserID,
) (*UserOrders, error) {

	spec := search.
		NewSpecificationTemplate().
		WithSearch(UserIDCondition(owner)).
		WithSearch(NotStatusesCondition([]OrderStatus{Registered})).
		WithOrder(search.AscOrder(CreatedAtOrderKey))

	orders, err := o.userOrdersRepository.FindBySpec(ctx, spec)
	if err != nil {
		return nil, err
	}

	return orders, nil

}

/* __________________________________________________ */

func (o *OrderManagerService) checkOrderExistence(
	ctx context.Context, owner common.UserID, order OrderID,
) error {

	ord, err := o.userOrderRepository.FindByID(ctx, order)
	if err != nil {
		return err
	}

	if ord == nil {
		return nil
	}

	if ord.Order.Identity().ID() == order.ID() &&
		ord.Owner.ID() == owner.ID() {
		return errx.NewErrorWithEntityCode(ErrorOrderEntity, ErrorOrderAlreadyLoaded)
	}

	if ord.Order.Identity().ID() == order.ID() &&
		ord.Owner.ID() != owner.ID() {
		return errx.NewErrorWithEntityCode(ErrorOrderEntity, ErrorOrderAlreadyExists)
	}

	return errx.NewErrorWithEntityCode(errx.ErrorInternalCode, errx.ErrorInternalMessage)

}

func (o *OrderManagerService) associateOrder(
	ctx context.Context, userID common.UserID, id OrderID,
) (*UserOrder, error) {
	userOrder := NewUserOrder(userID, *AsNewOrder(id))
	return o.userOrderRepository.NonAutoSave(ctx, *userOrder)
}

/* __________________________________________________ */
