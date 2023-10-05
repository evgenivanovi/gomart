package account

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/withdrawal"
	md "github.com/evgenivanovi/gomart/internal/model"
	withdraw_md "github.com/evgenivanovi/gomart/internal/model/withdraw"
)

/* __________________________________________________ */

type GetWithdrawalHistoryUsecase interface {
	Execute(context.Context, WithdrawalHistoryRequest) (WithdrawalHistoryResponse, error)
}

type GetWithdrawalHistoryUsecaseService struct {
	transactor        common.Transactor
	withdrawalManager withdrawal.WithdrawalManager
}

func ProvideGetWithdrawalHistoryUsecaseService(
	transactor common.Transactor,
	withdrawalManager withdrawal.WithdrawalManager,
) *GetWithdrawalHistoryUsecaseService {
	return &GetWithdrawalHistoryUsecaseService{
		transactor:        transactor,
		withdrawalManager: withdrawalManager,
	}
}

func (uc *GetWithdrawalHistoryUsecaseService) Execute(
	ctx context.Context, request WithdrawalHistoryRequest,
) (response WithdrawalHistoryResponse, err error) {

	user := common.NewUserID(request.Payload.UserID)
	history, err := uc.withdrawalManager.GetHistory(ctx, *user)

	if err != nil {
		return uc.toEmptyResponse(), err
	}

	return uc.toResponse(*history), nil

}

func (uc *GetWithdrawalHistoryUsecaseService) toEmptyResponse() WithdrawalHistoryResponse {
	return WithdrawalHistoryResponse{}
}

func (uc *GetWithdrawalHistoryUsecaseService) toResponse(
	response withdrawal.WithdrawalHistory,
) WithdrawalHistoryResponse {

	withdrawalModels := make([]withdraw_md.Withdrawal, 0)
	for _, elem := range response.Withdrawals {
		model := withdraw_md.Withdrawal{
			Order:  elem.Data().Order,
			Amount: elem.Data().Amount.Amount(),
			Metadata: md.Metadata{
				CreatedAt: elem.Metadata().CreatedAt,
				UpdatedAt: elem.Metadata().UpdatedAt,
				DeletedAt: elem.Metadata().DeletedAt,
			},
		}
		withdrawalModels = append(withdrawalModels, model)
	}

	withdrawalHistory := withdraw_md.WithdrawalHistory{
		UserID:      response.UserID.ID(),
		Withdrawals: withdrawalModels,
	}

	return WithdrawalHistoryResponse{
		Payload: WithdrawalHistoryResponsePayload{
			Withdrawals: withdrawalHistory,
		},
	}

}

/* __________________________________________________ */
