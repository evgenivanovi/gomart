package account

import (
	"github.com/evgenivanovi/gomart/internal/boot/common"
	"github.com/evgenivanovi/gomart/internal/boot/pg"
	balance_dm "github.com/evgenivanovi/gomart/internal/domain/balance"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/balance"
)

/* __________________________________________________ */

var (
	BalanceReadRepository  balance_dm.BalanceReadRepository
	BalanceWriteRepository balance_dm.BalanceWriteRepository
	BalanceRepository      balance_dm.BalanceRepository

	BalanceManager balance_dm.BalanceManager
)

/* __________________________________________________ */

func BootBalance() {

	BalanceReadRepository = balance.ProvideBalanceReadRepositoryService(
		pg.PostgresReadRequester,
	)

	BalanceWriteRepository = balance.ProvideBalanceWriteRepositoryService(
		pg.PostgresWriteRequester,
	)

	BalanceRepository = balance_dm.ProvideBalanceRepositoryService(
		BalanceReadRepository, BalanceWriteRepository,
	)

	BalanceManager = balance_dm.ProvideBalanceManagerService(
		common.Transactor,
		BalanceRepository,
	)

}

/* __________________________________________________ */
