package user

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/auth/user"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	pgcommon "github.com/evgenivanovi/gomart/internal/repository/postgres/common"
	"github.com/evgenivanovi/gomart/pkg/pg"
	"github.com/evgenivanovi/gomart/pkg/std"
)

/* __________________________________________________ */

type UserAuthWriteRepositoryService struct {
	requester pg.WriteRequester
}

func ProvideUserAuthWriteRepositoryService(
	requester pg.WriteRequester,
) *UserAuthWriteRepositoryService {
	return &UserAuthWriteRepositoryService{
		requester: requester,
	}
}

func (r *UserAuthWriteRepositoryService) AutoSave(
	ctx context.Context, data user.UserData, metadata core.Metadata,
) (*user.User, error) {

	var id int64
	command, args := insertOneStatement(FromUserData(data, metadata))

	err := r.requester.ExecuteReturningWithDefaultOnError(ctx, &id, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return user.NewUser(*common.NewUserID(id), data, metadata), nil

}

func (r *UserAuthWriteRepositoryService) AutoSaveAll(
	ctx context.Context, datas []std.Pair[user.UserData, core.Metadata],
) error {
	command, args := insertAllStatement(FromUsersData(datas))
	err := r.requester.ExecuteWithDefaultOnError(ctx, command, args)
	return r.translateError(err)
}

func (r *UserAuthWriteRepositoryService) translateError(err error) error {
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
