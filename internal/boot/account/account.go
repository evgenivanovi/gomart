package account

import (
	"github.com/evgenivanovi/gomart/internal/boot/common"
	"github.com/evgenivanovi/gomart/internal/usecase/account"
)

/* __________________________________________________ */

var (
	GetBalanceUseCase           account.GetBalanceUsecase
	CreateWithdrawUseCase       account.CreateWithdrawUsecase
	GetWithdrawalHistoryUseCase account.GetWithdrawalHistoryUsecase
)

/* __________________________________________________ */

func BootAccount() {

	GetBalanceUseCase = account.ProvideGetBalanceUsecaseService(
		BalanceManager, WithdrawManager,
	)

	CreateWithdrawUseCase = account.ProvideCreateWithdrawUsecaseService(
		common.Transactor,
		BalanceManager,
		WithdrawManager,
		WithdrawalManager,
	)

	GetWithdrawalHistoryUseCase = account.ProvideGetWithdrawalHistoryUsecaseService(
		common.Transactor,
		WithdrawalManager,
	)

}

/* __________________________________________________ */
