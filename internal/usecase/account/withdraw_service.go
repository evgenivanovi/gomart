package account

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/balance"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/withdraw"
	"github.com/evgenivanovi/gomart/internal/domain/withdrawal"
)

/* __________________________________________________ */

type CreateWithdrawUsecase interface {
	Execute(context.Context, WithdrawRequest) error
}

type CreateWithdrawUsecaseService struct {
	transactor        common.Transactor
	balanceManager    balance.BalanceManager
	withdrawManager   withdraw.WithdrawManager
	withdrawalManager withdrawal.WithdrawalManager
}

func ProvideCreateWithdrawUsecaseService(
	transactor common.Transactor,
	balanceManager balance.BalanceManager,
	withdrawManager withdraw.WithdrawManager,
	withdrawalManager withdrawal.WithdrawalManager,
) *CreateWithdrawUsecaseService {
	return &CreateWithdrawUsecaseService{
		transactor:        transactor,
		balanceManager:    balanceManager,
		withdrawManager:   withdrawManager,
		withdrawalManager: withdrawalManager,
	}
}

func (uc *CreateWithdrawUsecaseService) Execute(
	ctx context.Context, request WithdrawRequest,
) (err error) {

	ctx = uc.transactor.StartEx(ctx)
	defer uc.transactor.CloseEx(ctx, err)

	user := common.NewUserID(request.Payload.UserID)
	amount := common.NewPoints(request.Payload.Amount)
	withdrawalData := withdrawal.NewWithdrawalData(*user, *amount, request.Payload.Order)

	_, err = uc.balanceManager.Minus(ctx, *user, *amount)
	if err != nil {
		return err
	}

	_, err = uc.withdrawalManager.Create(ctx, *withdrawalData)
	if err != nil {
		return err
	}

	_, err = uc.withdrawManager.Add(ctx, *user, *amount)
	if err != nil {
		return err
	}

	return nil

}

/* __________________________________________________ */
