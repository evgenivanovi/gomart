package balance

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/balance"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	pgcommon "github.com/evgenivanovi/gomart/internal/repository/postgres/common"
	"github.com/evgenivanovi/gomart/pkg/pg"
)

/* __________________________________________________ */

type BalanceWriteRepositoryService struct {
	requester pg.WriteRequester
}

func ProvideBalanceWriteRepositoryService(
	requester pg.WriteRequester,
) *BalanceWriteRepositoryService {
	return &BalanceWriteRepositoryService{
		requester: requester,
	}
}

func (r *BalanceWriteRepositoryService) AutoSave(
	ctx context.Context, data balance.BalanceData,
) (*balance.Balance, error) {

	var id int64
	command, args := insertOneStatement(FromBalanceData(data))

	err := r.requester.ExecuteReturningWithDefaultOnError(ctx, &id, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return balance.NewBalance(*balance.NewBalanceID(id), data), nil

}

func (r *BalanceWriteRepositoryService) Update(
	ctx context.Context, data balance.Balance,
) (*balance.Balance, error) {

	command, args := updateOneStatement(FromBalance(data))

	err := r.requester.ExecuteWithDefaultOnError(ctx, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *BalanceWriteRepositoryService) translateError(err error) error {
	if err != nil {
		err = pgcommon.TranslateWriteError(err)
		err = pg.WithEntity(err, balance.ErrorBalanceEntity)
		if pg.ErrorCode(err) == core.ErrorNotFoundCode {
			return nil
		}
		return err
	}
	return nil
}

/* __________________________________________________ */
