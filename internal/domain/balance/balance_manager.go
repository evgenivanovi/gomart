package balance

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type BalanceManager interface {
	Get(ctx context.Context, user common.UserID) (*Balance, error)
	Create(ctx context.Context, user common.UserID) (*Balance, error)
	Add(ctx context.Context, user common.UserID, amount common.Points) (*Balance, error)
	Minus(ctx context.Context, user common.UserID, amount common.Points) (*Balance, error)
}

type BalanceManagerService struct {
	transactor common.Transactor
	repository BalanceRepository
}

func ProvideBalanceManagerService(
	transactor common.Transactor,
	repository BalanceRepository,
) *BalanceManagerService {
	return &BalanceManagerService{
		transactor: transactor,
		repository: repository,
	}
}

func (b *BalanceManagerService) Get(
	ctx context.Context, user common.UserID,
) (*Balance, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(UserIDCondition(user))
	return b.repository.FindOneBySpec(ctx, spec)
}

func (b *BalanceManagerService) Create(
	ctx context.Context, user common.UserID,
) (*Balance, error) {
	balance, err := b.repository.AutoSave(ctx, *NewEmptyBalanceData(user))
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (b *BalanceManagerService) Add(
	ctx context.Context, user common.UserID, amount common.Points,
) (*Balance, error) {

	spec := search.
		NewSpecificationTemplate().
		WithSearch(UserIDCondition(user))

	balance, err := b.repository.FindOneBySpecExclusively(ctx, spec)
	if err != nil {
		return nil, err
	}

	if balance == nil {
		return nil, errx.NewErrorWithEntityCode(ErrorBalanceEntity, core.ErrorNotFoundCode)
	}

	data := balance.Data()
	data.Balance = data.Balance.Add(amount.Amount())
	newBalance := balance.WithData(data)

	balance, err = b.repository.Update(ctx, newBalance)
	if err != nil {
		return nil, err
	}

	return balance, err

}

func (b *BalanceManagerService) Minus(
	ctx context.Context, user common.UserID, amount common.Points,
) (*Balance, error) {

	spec := search.
		NewSpecificationTemplate().
		WithSearch(UserIDCondition(user))

	balance, err := b.repository.FindOneBySpecExclusively(ctx, spec)
	if err != nil {
		return nil, err
	}

	if balance == nil {
		return nil, errx.NewErrorWithEntityCode(ErrorBalanceEntity, core.ErrorNotFoundCode)
	}

	if balance.Data().Balance.LessThan(amount.Amount()) {
		return nil, errx.NewErrorWithEntityCode(ErrorBalanceEntity, ErrorBalanceNotEnough)
	}

	data := balance.Data()
	data.Balance = data.Balance.Minus(amount.Amount())
	newBalance := balance.WithData(data)

	balance, err = b.repository.Update(ctx, newBalance)
	if err != nil {
		return nil, err
	}

	return balance, err

}

/* __________________________________________________ */
