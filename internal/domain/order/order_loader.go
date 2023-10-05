package order

import (
	"context"
	"time"

	"github.com/evgenivanovi/gomart/internal/domain/balance"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/pkg/search"
	"github.com/evgenivanovi/gomart/pkg/std"
	"github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
	"github.com/evgenivanovi/gomart/pkg/stdx/xsync"
)

/* __________________________________________________ */

type OrderLoadingSchedulerOptions struct {
	interval time.Duration
}

func NewOrderLoadingSchedulerOptions(
	interval time.Duration,
) *OrderLoadingSchedulerOptions {
	return &OrderLoadingSchedulerOptions{
		interval: interval,
	}
}

/* __________________________________________________ */

type OrderLoadingScheduler interface {
	Schedule()
	Close()
}

type OrderLoadingSchedulingService struct {
	scheduler *xsync.ScheduledExecutor
}

func ProvideOrderLoadingSchedulingService(
	schedulerOptions OrderLoadingSchedulerOptions,
	executorOptions OrderLoadingExecutorOptions,
	executor OrderLoadingExecutor,
) *OrderLoadingSchedulingService {

	scheduler := xsync.NewScheduledExecutor(
		schedulerOptions.interval,
		func() { executor.Execute(executorOptions) },
	)

	return &OrderLoadingSchedulingService{
		scheduler: scheduler,
	}

}

func (o *OrderLoadingSchedulingService) Schedule() {
	o.scheduler.Execute()
}

func (o *OrderLoadingSchedulingService) Close() {
	o.scheduler.Close()
}

/* __________________________________________________ */

type OrderLoadingExecutorOptions struct {
	OrderLoadSize int64
}

func NewOrderLoadingExecutorOptions(
	orderLoadSize int64,
) *OrderLoadingExecutorOptions {
	return &OrderLoadingExecutorOptions{
		OrderLoadSize: orderLoadSize,
	}
}

/* __________________________________________________ */

type OrderLoadingExecutor interface {
	Execute(options OrderLoadingExecutorOptions)
}

type OrderLoadingExecutorService struct {
	transactor           common.Transactor
	loyaltyOperations    OrderLoyaltyOperations
	userOrderRepository  UserOrderRepository
	userOrdersRepository UserOrdersRepository
	balanceManager       balance.BalanceManager
}

func ProvideOrderLoadingExecutorService(
	transactor common.Transactor,
	loyaltyOperations OrderLoyaltyOperations,
	userOrderRepository UserOrderRepository,
	userOrdersRepository UserOrdersRepository,
	balanceManager balance.BalanceManager,
) *OrderLoadingExecutorService {
	return &OrderLoadingExecutorService{
		transactor:           transactor,
		loyaltyOperations:    loyaltyOperations,
		userOrderRepository:  userOrderRepository,
		userOrdersRepository: userOrdersRepository,
		balanceManager:       balanceManager,
	}
}

func (o *OrderLoadingExecutorService) Execute(
	options OrderLoadingExecutorOptions,
) {

	err := o.transactor.Within(
		context.Background(),
		o.load(options),
	)

	if err != nil {
		slog.Log().Debug("order loading failed", slog.ErrAttr(err))
	}

}

func (o *OrderLoadingExecutorService) load(
	options OrderLoadingExecutorOptions,
) func(context.Context) error {

	var findOrdersAction = func(ctx context.Context) ([]*UserOrder, error) {

		spec := search.
			NewSpecificationTemplate().
			WithSearch(StatusesCondition([]OrderStatus{New, Registered})).
			WithSlice(search.WithLimit(options.OrderLoadSize))

		orders, err := o.userOrderRepository.FindManyBySpecExclusively(ctx, spec)
		if err != nil {
			return nil, err
		}

		return orders, nil

	}

	var getLoyaltyAction = func(ctx context.Context, order *UserOrder) (*OrderLoyalty, error) {
		return o.loyaltyOperations.Get(ctx, order.Order.Identity())
	}

	var updateOrderAction = func(ctx context.Context, order *UserOrder, loyalty *OrderLoyalty) error {

		userOrder := NewUserOrder(
			order.Owner,
			*NewOrder(order.Order.Identity(), loyalty.Data(), order.Order.Metadata()),
		)

		_, err := o.userOrderRepository.Update(ctx, *userOrder)
		if err != nil {
			return err
		}

		return nil

	}

	var updateBalanceAction = func(ctx context.Context, order *UserOrder, loyalty *OrderLoyalty) error {

		if std.IsPositiveFloat64(loyalty.Data().Accrual) {

			points := common.NewPoints(loyalty.Data().Accrual)
			_, err := o.balanceManager.Add(ctx, order.Owner, *points)

			if err != nil {
				return err
			}

		}

		return nil

	}

	return func(ctx context.Context) error {

		orders, orderErr := findOrdersAction(ctx)
		if orderErr != nil {
			return orderErr
		}

		for _, order := range orders {
			loyalty, loyaltyErr := getLoyaltyAction(ctx, order)
			if loyaltyErr != nil {
				return loyaltyErr
			}

			updateErr := updateOrderAction(ctx, order, loyalty)
			if updateErr != nil {
				return updateErr
			}

			updateBalanceErr := updateBalanceAction(ctx, order, loyalty)
			if updateBalanceErr != nil {
				return updateBalanceErr
			}
		}

		return nil

	}

}

/* __________________________________________________ */
