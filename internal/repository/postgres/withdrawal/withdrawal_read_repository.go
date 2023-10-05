package withdrawal

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/domain/withdraw"
	"github.com/evgenivanovi/gomart/internal/domain/withdrawal"
	pgcommon "github.com/evgenivanovi/gomart/internal/repository/postgres/common"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/table"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/evgenivanovi/gomart/pkg/pg"
	"github.com/evgenivanovi/gomart/pkg/search"
	pgsearch "github.com/evgenivanovi/gomart/pkg/search/jet/pg"
	pgjet "github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

type WithdrawalReadRepositoryService struct {
	requester     pg.ReadRequester
	searchMapping map[search.Key]pgjet.Column
	orderMapping  map[search.Key]pgjet.Column
}

func ProvideWithdrawalReadRepositoryService(
	requester pg.ReadRequester,
) *WithdrawalReadRepositoryService {

	searchMapping := make(map[search.Key]pgjet.Column)
	searchMapping[withdrawal.IDSearchKey] = table.Withdrawals.ID
	searchMapping[withdrawal.UserIDSearchKey] = table.Withdrawals.UserID

	orderMapping := make(map[search.Key]pgjet.Column)

	return &WithdrawalReadRepositoryService{
		requester:     requester,
		searchMapping: searchMapping,
		orderMapping:  orderMapping,
	}

}

func (r *WithdrawalReadRepositoryService) GetByID(
	ctx context.Context, id withdrawal.WithdrawalID,
) (*withdrawal.Withdrawal, error) {

	res, err := r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errx.NewErrorWithEntityCode(
			withdraw.ErrorWithdrawEntity,
			core.ErrorNotFoundCode,
		)
	}

	return res, nil

}

func (r *WithdrawalReadRepositoryService) FindByID(
	ctx context.Context, id withdrawal.WithdrawalID,
) (*withdrawal.Withdrawal, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(withdrawal.IdentityCondition(id))
	return r.FindOneBySpec(ctx, spec)
}

func (r *WithdrawalReadRepositoryService) FindByIDs(
	ctx context.Context, ids []withdrawal.WithdrawalID,
) ([]*withdrawal.Withdrawal, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(withdrawal.IdentitiesCondition(ids))
	return r.FindManyBySpec(ctx, spec)
}

func (r *WithdrawalReadRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*withdrawal.Withdrawal, error) {
	var dst model.Withdrawals
	query, args := r.query(spec, nil)

	err := r.requester.ExecuteOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToWithdrawal(&dst), nil
}

func (r *WithdrawalReadRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*withdrawal.Withdrawal, error) {
	var dst = make([]*model.Withdrawals, 0)
	query, args := r.query(spec, nil)

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToWithdrawals(dst), nil
}

func (r *WithdrawalReadRepositoryService) FindOneBySpecExclusively(
	ctx context.Context, spec search.Specification,
) (*withdrawal.Withdrawal, error) {
	var dst model.Withdrawals
	query, args := r.query(spec, pgjet.UPDATE())

	err := r.requester.ExecuteOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToWithdrawal(&dst), nil
}

func (r *WithdrawalReadRepositoryService) FindManyBySpecExclusively(
	ctx context.Context, spec search.Specification,
) ([]*withdrawal.Withdrawal, error) {
	var dst = make([]*model.Withdrawals, 0)
	query, args := r.query(spec, pgjet.UPDATE())

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToWithdrawals(dst), nil
}

func (r *WithdrawalReadRepositoryService) query(
	spec search.Specification,
	lock pgjet.RowLock,
) (string, []interface{}) {
	searchExp := pgsearch.SearchExpression(spec, r.searchMapping)
	orderExp := pgsearch.OrderExpression(spec, r.orderMapping)
	return buildQuery(searchExp, orderExp, lock, *spec.SliceConditions())
}

func (r *WithdrawalReadRepositoryService) translateError(err error) error {
	if err != nil {
		err = pgcommon.TranslateReadError(err)
		err = pg.WithEntity(err, withdraw.ErrorWithdrawEntity)
		if pg.ErrorCode(err) == core.ErrorNotFoundCode {
			return nil
		}
		return err
	}
	return nil
}

/* __________________________________________________ */
