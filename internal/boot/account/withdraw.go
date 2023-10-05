package account

import (
	"github.com/evgenivanovi/gomart/internal/boot/common"
	"github.com/evgenivanovi/gomart/internal/boot/pg"
	withdraw_dm "github.com/evgenivanovi/gomart/internal/domain/withdraw"
	withdrawal_dm "github.com/evgenivanovi/gomart/internal/domain/withdrawal"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/withdraw"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/withdrawal"
)

/* __________________________________________________ */

var (
	WithdrawReadRepository  withdraw_dm.WithdrawReadRepository
	WithdrawWriteRepository withdraw_dm.WithdrawWriteRepository
	WithdrawRepository      withdraw_dm.WithdrawRepository

	WithdrawManager withdraw_dm.WithdrawManager

	WithdrawalReadRepository  withdrawal_dm.WithdrawalReadRepository
	WithdrawalWriteRepository withdrawal_dm.WithdrawalWriteRepository
	WithdrawalRepository      withdrawal_dm.WithdrawalRepository

	WithdrawalManager withdrawal_dm.WithdrawalManager
)

/* __________________________________________________ */

func BootWithdraw() {

	WithdrawReadRepository = withdraw.ProvideWithdrawReadRepositoryService(
		pg.PostgresReadRequester,
	)

	WithdrawWriteRepository = withdraw.ProvideWithdrawWriteRepositoryService(
		pg.PostgresWriteRequester,
	)

	WithdrawRepository = withdraw_dm.ProvideWithdrawRepositoryService(
		WithdrawReadRepository, WithdrawWriteRepository,
	)

	WithdrawManager = withdraw_dm.ProvideWithdrawManagerService(
		common.Transactor, WithdrawRepository,
	)

	WithdrawalReadRepository = withdrawal.ProvideWithdrawalReadRepositoryService(
		pg.PostgresReadRequester,
	)

	WithdrawalWriteRepository = withdrawal.ProvideWithdrawalWriteRepositoryService(
		pg.PostgresWriteRequester,
	)

	WithdrawalRepository = withdrawal_dm.ProvideWithdrawalRepositoryService(
		WithdrawalReadRepository, WithdrawalWriteRepository,
	)

	WithdrawalManager = withdrawal_dm.ProvideWithdrawalManagerService(
		WithdrawalRepository,
	)

}

/* __________________________________________________ */
