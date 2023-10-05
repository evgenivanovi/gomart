package user

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/auth/user"
	"github.com/evgenivanovi/gomart/internal/domain/common"
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

type UserAuthReadRepositoryService struct {
	requester     pg.ReadRequester
	searchMapping map[search.Key]pgjet.Column
	orderMapping  map[search.Key]pgjet.Column
}

func ProvideUserAuthReadRepositoryService(
	requester pg.ReadRequester,
) *UserAuthReadRepositoryService {

	searchMapping := make(map[search.Key]pgjet.Column)
	searchMapping[user.IDSearchKey] = table.Users.ID
	searchMapping[user.UsernameSearchKey] = table.Users.Username

	orderMapping := make(map[search.Key]pgjet.Column)

	return &UserAuthReadRepositoryService{
		requester:     requester,
		searchMapping: searchMapping,
		orderMapping:  orderMapping,
	}

}

func (r *UserAuthReadRepositoryService) GetByID(
	ctx context.Context, id common.UserID,
) (*user.User, error) {

	res, err := r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errx.NewErrorWithEntityCode(
			user.ErrorUserEntity,
			core.ErrorNotFoundCode,
		)
	}

	return res, nil

}

func (r *UserAuthReadRepositoryService) FindByID(
	ctx context.Context, id common.UserID,
) (*user.User, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(user.IdentityCondition(id))
	return r.FindOneBySpec(ctx, spec)
}

func (r *UserAuthReadRepositoryService) FindByIDs(
	ctx context.Context, ids []common.UserID,
) ([]*user.User, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(user.IdentitiesCondition(ids))
	return r.FindManyBySpec(ctx, spec)
}

func (r *UserAuthReadRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*user.User, error) {
	var dst model.Users
	query, args := r.query(spec)

	err := r.requester.ExecuteOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToUser(&dst), nil
}

func (r *UserAuthReadRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*user.User, error) {
	var dst = make([]*model.Users, 0)
	query, args := r.query(spec)

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToUsers(dst), nil
}

func (r *UserAuthReadRepositoryService) query(
	spec search.Specification,
) (string, []interface{}) {
	searchExp := pgsearch.SearchExpression(spec, r.searchMapping)
	orderExp := pgsearch.OrderExpression(spec, r.orderMapping)
	return buildQuery(searchExp, orderExp, nil, *spec.SliceConditions())
}

func (r *UserAuthReadRepositoryService) translateError(err error) error {
	if err != nil {
		err = pgcommon.TranslateReadError(err)
		err = pg.WithEntity(err, user.ErrorUserEntity)
		if pg.ErrorCode(err) == core.ErrorNotFoundCode {
			return nil
		}
		return err
	}
	return nil
}

/* __________________________________________________ */
