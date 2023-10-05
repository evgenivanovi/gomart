package withdraw

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type WithdrawManager interface {
	Get(ctx context.Context, user common.UserID) (*Withdraw, error)
	Create(ctx context.Context, user common.UserID) (*Withdraw, error)
	Add(ctx context.Context, user common.UserID, amount common.Points) (*Withdraw, error)
	Minus(ctx context.Context, user common.UserID, amount common.Points) (*Withdraw, error)
}

type WithdrawManagerService struct {
	transactor common.Transactor
	repository WithdrawRepository
}

func ProvideWithdrawManagerService(
	transactor common.Transactor,
	repository WithdrawRepository,
) *WithdrawManagerService {
	return &WithdrawManagerService{
		transactor: transactor,
		repository: repository,
	}
}

func (b *WithdrawManagerService) Get(
	ctx context.Context, user common.UserID,
) (*Withdraw, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(UserIDCondition(user))
	return b.repository.FindOneBySpec(ctx, spec)
}

func (b *WithdrawManagerService) Create(
	ctx context.Context, user common.UserID,
) (*Withdraw, error) {
	balance, err := b.repository.AutoSave(ctx, *NewEmptyWithdrawData(user))
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (b *WithdrawManagerService) Add(
	ctx context.Context, user common.UserID, amount common.Points,
) (*Withdraw, error) {

	spec := search.
		NewSpecificationTemplate().
		WithSearch(UserIDCondition(user))

	withdraw, err := b.repository.FindOneBySpecExclusively(ctx, spec)
	if err != nil {
		return nil, err
	}

	if withdraw == nil {
		return nil, errx.NewErrorWithEntityCode(ErrorWithdrawEntity, core.ErrorNotFoundCode)
	}

	data := withdraw.Data()
	data.Withdraw = data.Withdraw.Add(amount.Amount())
	newWithdraw := withdraw.WithData(data)

	result, err := b.repository.Update(ctx, newWithdraw)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (b *WithdrawManagerService) Minus(
	ctx context.Context, user common.UserID, amount common.Points,
) (*Withdraw, error) {

	spec := search.
		NewSpecificationTemplate().
		WithSearch(UserIDCondition(user))

	withdraw, err := b.repository.FindOneBySpecExclusively(ctx, spec)
	if err != nil {
		return nil, err
	}

	if withdraw == nil {
		return nil, errx.NewErrorWithEntityCode(ErrorWithdrawEntity, core.ErrorNotFoundCode)
	}

	data := withdraw.Data()
	data.Withdraw = data.Withdraw.Minus(amount.Amount())
	newBalance := withdraw.WithData(data)

	result, err := b.repository.Update(ctx, newBalance)
	if err != nil {
		return nil, err
	}

	return result, nil

}

/* __________________________________________________ */
