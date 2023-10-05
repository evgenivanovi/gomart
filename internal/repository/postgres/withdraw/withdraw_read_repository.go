package withdraw

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/domain/withdraw"
	pgcommon "github.com/evgenivanovi/gomart/internal/repository/postgres/common"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/evgenivanovi/gomart/pkg/pg"
	"github.com/evgenivanovi/gomart/pkg/search"
	pgsearch "github.com/evgenivanovi/gomart/pkg/search/jet/pg"
	pgjet "github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

type WithdrawReadRepositoryService struct {
	requester     pg.ReadRequester
	searchMapping map[search.Key]pgjet.Column
	orderMapping  map[search.Key]pgjet.Column
}

func ProvideWithdrawReadRepositoryService(
	requester pg.ReadRequester,
) *WithdrawReadRepositoryService {

	searchMapping := make(map[search.Key]pgjet.Column)
	orderMapping := make(map[search.Key]pgjet.Column)

	return &WithdrawReadRepositoryService{
		requester:     requester,
		searchMapping: searchMapping,
		orderMapping:  orderMapping,
	}

}

func (r *WithdrawReadRepositoryService) GetByID(
	ctx context.Context, id withdraw.WithdrawID,
) (*withdraw.Withdraw, error) {

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

func (r *WithdrawReadRepositoryService) FindByID(
	ctx context.Context, id withdraw.WithdrawID,
) (*withdraw.Withdraw, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(withdraw.IdentityCondition(id))
	return r.FindOneBySpec(ctx, spec)
}

func (r *WithdrawReadRepositoryService) FindByIDs(
	ctx context.Context, ids []withdraw.WithdrawID,
) ([]*withdraw.Withdraw, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(withdraw.IdentitiesCondition(ids))
	return r.FindManyBySpec(ctx, spec)
}

func (r *WithdrawReadRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*withdraw.Withdraw, error) {
	var dst model.Withdraws
	query, args := r.query(spec, nil)

	err := r.requester.ExecuteOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToWithdraw(&dst), nil
}

func (r *WithdrawReadRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*withdraw.Withdraw, error) {
	var dst = make([]*model.Withdraws, 0)
	query, args := r.query(spec, nil)

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToWithdraws(dst), nil
}

func (r *WithdrawReadRepositoryService) FindOneBySpecExclusively(
	ctx context.Context, spec search.Specification,
) (*withdraw.Withdraw, error) {
	var dst model.Withdraws
	query, args := r.query(spec, pgjet.UPDATE())

	err := r.requester.ExecuteOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToWithdraw(&dst), nil
}

func (r *WithdrawReadRepositoryService) FindManyBySpecExclusively(
	ctx context.Context, spec search.Specification,
) ([]*withdraw.Withdraw, error) {
	var dst = make([]*model.Withdraws, 0)
	query, args := r.query(spec, pgjet.UPDATE())

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToWithdraws(dst), nil
}

func (r *WithdrawReadRepositoryService) query(
	spec search.Specification,
	lock pgjet.RowLock,
) (string, []interface{}) {
	searchExp := pgsearch.SearchExpression(spec, r.searchMapping)
	orderExp := pgsearch.OrderExpression(spec, r.orderMapping)
	return buildQuery(searchExp, orderExp, lock, *spec.SliceConditions())
}

func (r *WithdrawReadRepositoryService) translateError(err error) error {
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
