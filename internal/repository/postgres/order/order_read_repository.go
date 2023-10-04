package user

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/domain/order"
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

type UserOrderReadRepositoryService struct {
	requester     pg.ReadRequester
	searchMapping map[search.Key]pgjet.Column
	orderMapping  map[search.Key]pgjet.Column
}

func ProvideUserOrderReadRepositoryService(
	requester pg.ReadRequester,
) *UserOrderReadRepositoryService {

	searchMapping := make(map[search.Key]pgjet.Column)
	searchMapping[order.IDSearchKey] = table.Orders.ID
	searchMapping[order.UserIDSearchKey] = table.Orders.UserID
	searchMapping[order.StatusSearchKey] = table.Orders.Status

	orderMapping := make(map[search.Key]pgjet.Column)
	orderMapping[order.CreatedAtOrderKey] = table.Orders.CreatedAt
	orderMapping[order.UpdatedAtIDOrderKey] = table.Orders.UpdatedAt
	orderMapping[order.DeletedAtIDOrderKey] = table.Orders.DeletedAt

	return &UserOrderReadRepositoryService{
		requester:     requester,
		searchMapping: searchMapping,
		orderMapping:  orderMapping,
	}

}

func (r *UserOrderReadRepositoryService) GetByID(
	ctx context.Context, id order.OrderID,
) (*order.UserOrder, error) {

	res, err := r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errx.NewErrorWithEntityCode(
			order.ErrorOrderEntity,
			core.ErrorNotFoundCode,
		)
	}

	return res, nil

}

func (r *UserOrderReadRepositoryService) FindByID(
	ctx context.Context, id order.OrderID,
) (*order.UserOrder, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(order.IdentityCondition(id))
	return r.FindOneBySpec(ctx, spec)
}

func (r *UserOrderReadRepositoryService) FindByIDs(
	ctx context.Context, ids []order.OrderID,
) ([]*order.UserOrder, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(order.IdentitiesCondition(ids))
	return r.FindManyBySpec(ctx, spec)
}

func (r *UserOrderReadRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*order.UserOrder, error) {
	var dst model.Orders
	query, args := r.query(spec, nil)

	err := r.requester.ExecuteOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToUserOrder(&dst), nil
}

func (r *UserOrderReadRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*order.UserOrder, error) {
	var dst = make([]*model.Orders, 0)
	query, args := r.query(spec, nil)

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToUserOrderSlice(dst), nil
}

func (r *UserOrderReadRepositoryService) FindOneBySpecExclusively(
	ctx context.Context, spec search.Specification,
) (*order.UserOrder, error) {
	var dst model.Orders
	query, args := r.query(spec, pgjet.UPDATE())

	err := r.requester.ExecuteOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToUserOrder(&dst), nil
}

func (r *UserOrderReadRepositoryService) FindManyBySpecExclusively(
	ctx context.Context, spec search.Specification,
) ([]*order.UserOrder, error) {
	var dst = make([]*model.Orders, 0)
	query, args := r.query(spec, pgjet.UPDATE())

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToUserOrderSlice(dst), nil
}

func (r *UserOrderReadRepositoryService) query(
	spec search.Specification,
	lock pgjet.RowLock,
) (string, []interface{}) {
	searchExp := pgsearch.SearchExpression(spec, r.searchMapping)
	orderExp := pgsearch.OrderExpression(spec, r.orderMapping)
	return buildQuery(searchExp, orderExp, lock, *spec.SliceConditions())
}

func (r *UserOrderReadRepositoryService) translateError(err error) error {
	if err != nil {
		err = pgcommon.TranslateReadError(err)
		err = pg.WithEntity(err, order.ErrorOrderEntity)
		if pg.ErrorCode(err) == core.ErrorNotFoundCode {
			return nil
		}
		return err
	}
	return nil
}

/* __________________________________________________ */

type UserOrdersReadRepositoryService struct {
	requester     pg.ReadRequester
	searchMapping map[search.Key]pgjet.Column
	orderMapping  map[search.Key]pgjet.Column
}

func ProvideUserOrdersReadRepositoryService(
	requester pg.ReadRequester,
) *UserOrdersReadRepositoryService {

	searchMapping := make(map[search.Key]pgjet.Column)
	searchMapping[order.IDSearchKey] = table.Orders.ID
	searchMapping[order.UserIDSearchKey] = table.Orders.UserID
	searchMapping[order.StatusSearchKey] = table.Orders.Status

	orderMapping := make(map[search.Key]pgjet.Column)
	orderMapping[order.CreatedAtOrderKey] = table.Orders.CreatedAt
	orderMapping[order.UpdatedAtIDOrderKey] = table.Orders.UpdatedAt
	orderMapping[order.DeletedAtIDOrderKey] = table.Orders.DeletedAt

	return &UserOrdersReadRepositoryService{
		requester:     requester,
		searchMapping: searchMapping,
		orderMapping:  orderMapping,
	}

}

func (r *UserOrdersReadRepositoryService) GetByID(
	ctx context.Context, id common.UserID,
) (*order.UserOrders, error) {

	res, err := r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errx.NewErrorWithEntityCode(
			order.ErrorOrderEntity,
			core.ErrorNotFoundCode,
		)
	}

	return res, nil

}

func (r *UserOrdersReadRepositoryService) FindByID(
	ctx context.Context, id common.UserID,
) (*order.UserOrders, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(order.UserIDCondition(id))
	return r.FindBySpec(ctx, spec)
}

func (r *UserOrdersReadRepositoryService) FindBySpec(
	ctx context.Context, spec search.Specification,
) (*order.UserOrders, error) {
	var dst = make([]*model.Orders, 0)
	query, args := r.query(spec)

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToUserOrders(dst)
}

func (r *UserOrdersReadRepositoryService) query(
	spec search.Specification,
) (string, []interface{}) {
	searchExp := pgsearch.SearchExpression(spec, r.searchMapping)
	orderExp := pgsearch.OrderExpression(spec, r.orderMapping)
	return buildQuery(searchExp, orderExp, nil, *spec.SliceConditions())
}

func (r *UserOrdersReadRepositoryService) translateError(err error) error {
	if err != nil {
		err = pgcommon.TranslateReadError(err)
		err = pg.WithEntity(err, order.ErrorOrderEntity)
		if pg.ErrorCode(err) == core.ErrorNotFoundCode {
			return nil
		}
		return err
	}
	return nil
}

/* __________________________________________________ */
