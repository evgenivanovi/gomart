package balance

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/balance"
	"github.com/evgenivanovi/gomart/internal/domain/core"
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

type BalanceReadRepositoryService struct {
	requester     pg.ReadRequester
	searchMapping map[search.Key]pgjet.Column
	orderMapping  map[search.Key]pgjet.Column
}

func ProvideBalanceReadRepositoryService(
	requester pg.ReadRequester,
) *BalanceReadRepositoryService {

	searchMapping := make(map[search.Key]pgjet.Column)
	searchMapping[balance.IDSearchKey] = table.Balances.ID
	searchMapping[balance.UserIDSearchKey] = table.Balances.UserID

	orderMapping := make(map[search.Key]pgjet.Column)

	return &BalanceReadRepositoryService{
		requester:     requester,
		searchMapping: searchMapping,
		orderMapping:  orderMapping,
	}

}

func (r *BalanceReadRepositoryService) GetByID(
	ctx context.Context, id balance.BalanceID,
) (*balance.Balance, error) {

	res, err := r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errx.NewErrorWithEntityCode(
			balance.ErrorBalanceEntity,
			core.ErrorNotFoundCode,
		)
	}

	return res, nil

}

func (r *BalanceReadRepositoryService) FindByID(
	ctx context.Context, id balance.BalanceID,
) (*balance.Balance, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(balance.IdentityCondition(id))
	return r.FindOneBySpec(ctx, spec)
}

func (r *BalanceReadRepositoryService) FindByIDs(
	ctx context.Context, ids []balance.BalanceID,
) ([]*balance.Balance, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(balance.IdentitiesCondition(ids))
	return r.FindManyBySpec(ctx, spec)
}

func (r *BalanceReadRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*balance.Balance, error) {
	var dst model.Balances
	query, args := r.query(spec, nil)

	err := r.requester.ExecuteOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToBalance(&dst), nil
}

func (r *BalanceReadRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*balance.Balance, error) {
	var dst = make([]*model.Balances, 0)
	query, args := r.query(spec, nil)

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToBalances(dst), nil
}

func (r *BalanceReadRepositoryService) FindOneBySpecExclusively(
	ctx context.Context, spec search.Specification,
) (*balance.Balance, error) {
	var dst model.Balances
	query, args := r.query(spec, pgjet.UPDATE())

	err := r.requester.ExecuteOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToBalance(&dst), nil
}

func (r *BalanceReadRepositoryService) FindManyBySpecExclusively(
	ctx context.Context, spec search.Specification,
) ([]*balance.Balance, error) {
	var dst = make([]*model.Balances, 0)
	query, args := r.query(spec, pgjet.UPDATE())

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToBalances(dst), nil
}

func (r *BalanceReadRepositoryService) query(
	spec search.Specification,
	lock pgjet.RowLock,
) (string, []interface{}) {
	searchExp := pgsearch.SearchExpression(spec, r.searchMapping)
	orderExp := pgsearch.OrderExpression(spec, r.orderMapping)
	return buildQuery(searchExp, orderExp, lock, *spec.SliceConditions())
}

func (r *BalanceReadRepositoryService) translateError(err error) error {
	if err != nil {
		err = pgcommon.TranslateReadError(err)
		err = pg.WithEntity(err, balance.ErrorBalanceEntity)
		if pg.ErrorCode(err) == core.ErrorNotFoundCode {
			return nil
		}
		return err
	}
	return nil
}

/* __________________________________________________ */
