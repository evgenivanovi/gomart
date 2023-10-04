package order

import (
	"time"

	"github.com/evgenivanovi/gomart/internal/boot/account"
	"github.com/evgenivanovi/gomart/internal/boot/app"
	"github.com/evgenivanovi/gomart/internal/boot/common"
	"github.com/evgenivanovi/gomart/internal/boot/pg"
	order_dm "github.com/evgenivanovi/gomart/internal/domain/order"
	order_pg "github.com/evgenivanovi/gomart/internal/repository/postgres/order"
	order_uc "github.com/evgenivanovi/gomart/internal/usecase/order"
	"github.com/evgenivanovi/gomart/pkg/valid"
)

/* __________________________________________________ */

var (
	OrderLoadSize     = int64(10)
	OrderLoadInterval = time.Duration(10) * time.Second
)

var (
	UserOrderReadRepository  order_dm.UserOrderReadRepository
	UserOrderWriteRepository order_dm.UserOrderWriteRepository
	UserOrderRepository      order_dm.UserOrderRepository

	UserOrdersReadRepository order_dm.UserOrdersReadRepository
	UserOrdersRepository     order_dm.UserOrdersRepository

	OrderLoyaltyOperations order_dm.OrderLoyaltyOperations

	OrderManager          order_dm.OrderManager
	OrderValidator        order_dm.OrderValidator
	OrderLoadingExecutor  order_dm.OrderLoadingExecutor
	OrderLoadingScheduler order_dm.OrderLoadingScheduler

	OrderLoadUsecase order_uc.OrderLoadUsecase
	GetOrdersUsecase order_uc.GetOrdersUsecase
)

/* __________________________________________________ */

func BootOrder() {

	UserOrderReadRepository = order_pg.ProvideUserOrderReadRepositoryService(
		pg.PostgresReadRequester,
	)
	UserOrderWriteRepository = order_pg.ProvideUserOrderWriteRepositoryService(
		pg.PostgresWriteRequester,
	)
	UserOrderRepository = order_dm.ProvideUserOrderRepositoryService(
		UserOrderReadRepository, UserOrderWriteRepository,
	)

	UserOrdersReadRepository = order_pg.ProvideUserOrdersReadRepositoryService(
		pg.PostgresReadRequester,
	)
	UserOrdersRepository = order_dm.ProvideUserOrdersRepositoryService(
		UserOrdersReadRepository,
	)

	OrderValidator = order_dm.ProvideOrderValidatorService(valid.Luhn)

	OrderManager = order_dm.ProvideOrderManagerService(
		OrderLoyaltyOperations,
		OrderValidator,
		UserOrderRepository,
		UserOrdersRepository,
	)

	OrderLoadingExecutor = order_dm.ProvideOrderLoadingExecutorService(
		common.Transactor,
		OrderLoyaltyOperations,
		UserOrderRepository,
		UserOrdersRepository,
		account.BalanceManager,
	)

	OrderLoadingScheduler = order_dm.ProvideOrderLoadingSchedulingService(
		*order_dm.NewOrderLoadingSchedulerOptions(OrderLoadInterval),
		*order_dm.NewOrderLoadingExecutorOptions(OrderLoadSize),
		OrderLoadingExecutor,
	)

	app.Application.RegisterOnStartBackground(OrderLoadingScheduler.Schedule)
	app.Application.RegisterOnCloseBackground(OrderLoadingScheduler.Close)

	OrderLoadUsecase = order_uc.ProvideOrderLoadUsecaseService(
		common.Transactor, OrderManager,
	)

	GetOrdersUsecase = order_uc.ProvideGetOrderUsecaseService(
		common.Transactor, OrderManager,
	)

}

/* __________________________________________________ */
