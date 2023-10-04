package account

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/balance"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/withdraw"
)

/* __________________________________________________ */

type GetBalanceUsecase interface {
	Execute(context.Context, GetBalanceRequest) (GetBalanceResponse, error)
}

type GetBalanceUsecaseService struct {
	balanceManager  balance.BalanceManager
	withdrawManager withdraw.WithdrawManager
}

func ProvideGetBalanceUsecaseService(
	balanceManager balance.BalanceManager,
	withdrawManager withdraw.WithdrawManager,
) *GetBalanceUsecaseService {
	return &GetBalanceUsecaseService{
		balanceManager:  balanceManager,
		withdrawManager: withdrawManager,
	}
}

func (uc *GetBalanceUsecaseService) Execute(
	ctx context.Context, request GetBalanceRequest,
) (GetBalanceResponse, error) {

	user := common.NewUserID(request.Payload.UserID)

	balance, err := uc.balanceManager.Get(ctx, *user)
	if err != nil {
		return GetBalanceResponse{}, err
	}

	withdraw, err := uc.withdrawManager.Get(ctx, *user)
	if err != nil {
		return GetBalanceResponse{}, err
	}

	return GetBalanceResponse{
		Payload: GetBalanceResponsePayload{
			Balance:   balance.Data().Balance.Amount(),
			Withdrawn: withdraw.Data().Withdraw.Amount(),
		},
	}, nil

}

/* __________________________________________________ */
