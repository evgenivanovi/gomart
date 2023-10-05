package user

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/auth/user"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/domain/order"
	pgcommon "github.com/evgenivanovi/gomart/internal/repository/postgres/common"
	"github.com/evgenivanovi/gomart/pkg/pg"
)

/* __________________________________________________ */

type UserOrderWriteRepositoryService struct {
	requester pg.WriteRequester
}

func ProvideUserOrderWriteRepositoryService(
	requester pg.WriteRequester,
) *UserOrderWriteRepositoryService {
	return &UserOrderWriteRepositoryService{
		requester: requester,
	}
}

func (r *UserOrderWriteRepositoryService) NonAutoSave(
	ctx context.Context, data order.UserOrder,
) (*order.UserOrder, error) {

	command, args := insertOneStatement(FromUserOrder(data))

	err := r.requester.ExecuteWithDefaultOnError(ctx, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *UserOrderWriteRepositoryService) NonAutoSaveAll(
	ctx context.Context, datas []order.UserOrder,
) error {
	command, args := insertAllStatement(FromUserOrdersSlice(datas))
	err := r.requester.ExecuteWithDefaultOnError(ctx, command, args)
	return r.translateError(err)
}

func (r *UserOrderWriteRepositoryService) Update(
	ctx context.Context, data order.UserOrder,
) (*order.UserOrder, error) {

	command, args := updateOneStatement(FromUserOrder(data))

	err := r.requester.ExecuteWithDefaultOnError(ctx, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *UserOrderWriteRepositoryService) translateError(err error) error {
	if err != nil {
		err = pgcommon.TranslateWriteError(err)
		err = pg.WithEntity(err, user.ErrorUserEntity)
		if pg.ErrorCode(err) == core.ErrorNotFoundCode {
			return nil
		}
		return err
	}
	return nil
}

/* __________________________________________________ */
