package session

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/auth/session"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	pgcommon "github.com/evgenivanovi/gomart/internal/repository/postgres/common"
	"github.com/evgenivanovi/gomart/pkg/pg"
)

/* __________________________________________________ */

type SessionWriteRepositoryService struct {
	requester pg.WriteRequester
}

func ProvideSessionWriteRepositoryService(
	requester pg.WriteRequester,
) *SessionWriteRepositoryService {
	return &SessionWriteRepositoryService{
		requester: requester,
	}
}

func (r *SessionWriteRepositoryService) NonAutoSave(
	ctx context.Context, data session.Session,
) (*session.Session, error) {

	command, args := insertOneStatement(FromSession(data))

	err := r.requester.ExecuteWithDefaultOnError(ctx, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *SessionWriteRepositoryService) NonAutoSaveAll(
	ctx context.Context, data []session.Session,
) error {
	command, args := insertAllStatement(FromSessions(data))
	err := r.requester.ExecuteWithDefaultOnError(ctx, command, args...)
	return r.translateError(err)
}

func (r *SessionWriteRepositoryService) translateError(err error) error {
	if err != nil {
		err = pgcommon.TranslateWriteError(err)
		err = pg.WithEntity(err, session.ErrorSessionEntity)
		if pg.ErrorCode(err) == core.ErrorNotFoundCode {
			return nil
		}
		return err
	}
	return nil
}

/* __________________________________________________ */
