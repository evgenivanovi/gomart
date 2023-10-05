package account

import (
	"github.com/evgenivanovi/gomart/internal/http/infra/account"
)

/* __________________________________________________ */

var (
	GetBalanceController        *account.GetBalanceController
	WithdrawController          *account.WithdrawController
	WithdrawalHistoryController *account.WithdrawalHistoryController
)

/* __________________________________________________ */

func BootAccountHandlers() {

	GetBalanceController = account.ProvideGetBalanceController(
		GetBalanceUseCase,
	)

	WithdrawController = account.ProvideWithdrawController(
		CreateWithdrawUseCase,
	)

	WithdrawalHistoryController = account.ProvideWithdrawalHistoryController(
		GetWithdrawalHistoryUseCase,
	)

}

/* __________________________________________________ */
