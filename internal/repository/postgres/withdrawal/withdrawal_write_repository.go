package withdrawal

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/domain/withdraw"
	"github.com/evgenivanovi/gomart/internal/domain/withdrawal"
	pgcommon "github.com/evgenivanovi/gomart/internal/repository/postgres/common"
	"github.com/evgenivanovi/gomart/pkg/pg"
)

/* __________________________________________________ */

type WithdrawalWriteRepositoryService struct {
	requester pg.WriteRequester
}

func ProvideWithdrawalWriteRepositoryService(
	requester pg.WriteRequester,
) *WithdrawalWriteRepositoryService {
	return &WithdrawalWriteRepositoryService{
		requester: requester,
	}
}

func (r *WithdrawalWriteRepositoryService) AutoSave(
	ctx context.Context, data withdrawal.WithdrawalData, metadata core.Metadata,
) (*withdrawal.Withdrawal, error) {

	var id int64
	command, args := insertOneStatement(FromWithdrawalData(data, metadata))

	err := r.requester.ExecuteReturningWithDefaultOnError(ctx, &id, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return withdrawal.NewWithdrawal(*withdrawal.NewWithdrawalID(id), data, metadata), nil

}

func (r *WithdrawalWriteRepositoryService) translateError(err error) error {
	if err != nil {
		err = pgcommon.TranslateWriteError(err)
		err = pg.WithEntity(err, withdraw.ErrorWithdrawEntity)
		if pg.ErrorCode(err) == core.ErrorNotFoundCode {
			return nil
		}
		return err
	}
	return nil
}

/* __________________________________________________ */
