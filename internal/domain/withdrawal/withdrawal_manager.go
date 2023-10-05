package withdrawal

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type WithdrawalManager interface {
	Create(ctx context.Context, data WithdrawalData) (*Withdrawal, error)
	GetHistory(ctx context.Context, user common.UserID) (*WithdrawalHistory, error)
}

type WithdrawalManagerService struct {
	repository WithdrawalRepository
}

func ProvideWithdrawalManagerService(
	repository WithdrawalRepository,
) *WithdrawalManagerService {
	return &WithdrawalManagerService{
		repository: repository,
	}
}

func (w *WithdrawalManagerService) Create(
	ctx context.Context, data WithdrawalData,
) (*Withdrawal, error) {
	withdrawal, err := w.repository.AutoSave(ctx, data, *core.NewNowMetadata())
	if err != nil {
		return nil, err
	}
	return withdrawal, nil
}

func (w *WithdrawalManagerService) GetHistory(
	ctx context.Context, user common.UserID,
) (*WithdrawalHistory, error) {

	spec := search.
		NewSpecificationTemplate().
		WithSearch(UserIDCondition(user))

	withdrawals, err := w.repository.FindManyBySpec(ctx, spec)
	if err != nil {
		return nil, err
	}

	return NewWithdrawalHistory(user, ToWithdrawalValues(withdrawals)), nil

}

/* __________________________________________________ */
