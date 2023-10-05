package session

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/auth/session"
	"github.com/evgenivanovi/gomart/internal/domain/auth/user"
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

type SessionReadRepositoryService struct {
	requester     pg.ReadRequester
	searchMapping map[search.Key]pgjet.Column
	orderMapping  map[search.Key]pgjet.Column
}

func ProvideSessionReadRepositoryService(
	requester pg.ReadRequester,
) *SessionReadRepositoryService {

	searchMapping := make(map[search.Key]pgjet.Column)
	searchMapping[user.IDSearchKey] = table.Sessions.ID
	searchMapping[session.UserIDSearchKey] = table.Sessions.UserID

	orderMapping := make(map[search.Key]pgjet.Column)

	return &SessionReadRepositoryService{
		requester:     requester,
		searchMapping: searchMapping,
		orderMapping:  orderMapping,
	}

}

func (r *SessionReadRepositoryService) GetByID(
	ctx context.Context, id session.SessionID,
) (*session.Session, error) {

	res, err := r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errx.NewErrorWithEntityCode(
			session.ErrorSessionEntity,
			core.ErrorNotFoundCode,
		)
	}

	return res, nil

}

func (r *SessionReadRepositoryService) FindByID(
	ctx context.Context, id session.SessionID,
) (*session.Session, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(session.IdentityCondition(id))
	return r.FindOneBySpec(ctx, spec)
}

func (r *SessionReadRepositoryService) FindByIDs(
	ctx context.Context, ids []session.SessionID,
) ([]*session.Session, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(session.IdentitiesCondition(ids))
	return r.FindManyBySpec(ctx, spec)
}

func (r *SessionReadRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*session.Session, error) {
	var dst model.Sessions
	query, args := r.query(spec)

	err := r.requester.ExecuteOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToSession(&dst), nil
}

func (r *SessionReadRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*session.Session, error) {
	var dst = make([]*model.Sessions, 0)
	query, args := r.query(spec)

	err := r.requester.ExecuteManyWithScan(ctx, scanManyFunc(&dst), query, args)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToSessions(dst), nil
}

func (r *SessionReadRepositoryService) query(
	spec search.Specification,
) (string, []interface{}) {
	searchExp := pgsearch.SearchExpression(spec, r.searchMapping)
	orderExp := pgsearch.OrderExpression(spec, r.orderMapping)
	return buildQuery(searchExp, orderExp, nil, *spec.SliceConditions())
}

func (r *SessionReadRepositoryService) translateError(err error) error {
	if err != nil {
		err = pgcommon.TranslateReadError(err)
		err = pg.WithEntity(err, session.ErrorSessionEntity)
		if pg.ErrorCode(err) == core.ErrorNotFoundCode {
			return nil
		}
		return err
	}
	return nil
}

/* __________________________________________________ */
