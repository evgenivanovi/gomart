package withdraw

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/domain/withdraw"
	pgcommon "github.com/evgenivanovi/gomart/internal/repository/postgres/common"
	"github.com/evgenivanovi/gomart/pkg/pg"
)

/* __________________________________________________ */

type WithdrawWriteRepositoryService struct {
	requester pg.WriteRequester
}

func ProvideWithdrawWriteRepositoryService(
	requester pg.WriteRequester,
) *WithdrawWriteRepositoryService {
	return &WithdrawWriteRepositoryService{
		requester: requester,
	}
}

func (r *WithdrawWriteRepositoryService) AutoSave(
	ctx context.Context, data withdraw.WithdrawData,
) (*withdraw.Withdraw, error) {

	var id int64
	command, args := insertOneStatement(FromWithdrawData(data))

	err := r.requester.ExecuteReturningWithDefaultOnError(ctx, &id, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return withdraw.NewWithdraw(*withdraw.NewWithdrawID(id), data), nil

}

func (r *WithdrawWriteRepositoryService) Update(
	ctx context.Context, data withdraw.Withdraw,
) (*withdraw.Withdraw, error) {

	command, args := updateOneStatement(FromWithdraw(data))

	err := r.requester.ExecuteWithDefaultOnError(ctx, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *WithdrawWriteRepositoryService) translateError(err error) error {
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
