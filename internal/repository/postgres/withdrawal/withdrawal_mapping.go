package withdrawal

import (
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/domain/withdrawal"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/evgenivanovi/gomart/pkg/std"
)

/* __________________________________________________ */

func ToWithdrawal(record *model.Withdrawals) *withdrawal.Withdrawal {

	id := withdrawal.NewWithdrawalID(record.ID)

	data := withdrawal.NewWithdrawalData(
		*common.NewUserID(record.UserID),
		*common.NewPoints(record.Amount),
		record.Order,
	)

	metadata := core.NewMetadata(
		record.CreatedAt, record.UpdatedAt, record.DeletedAt,
	)

	return withdrawal.NewWithdrawal(*id, *data, *metadata)

}

func FromWithdrawal(entity withdrawal.Withdrawal) model.Withdrawals {
	return model.Withdrawals{
		ID: entity.Identity().ID(),
		// DATA
		UserID: entity.Data().UserID.ID(),
		Amount: entity.Data().Amount.Amount(),
		Order:  entity.Data().Order,
		// METADATA
		CreatedAt: entity.Metadata().CreatedAt,
		UpdatedAt: entity.Metadata().UpdatedAt,
		DeletedAt: entity.Metadata().DeletedAt,
	}
}

func FromWithdrawalData(data withdrawal.WithdrawalData, metadata core.Metadata) model.Withdrawals {
	return model.Withdrawals{
		// DATA
		UserID: data.UserID.ID(),
		Amount: data.Amount.Amount(),
		Order:  data.Order,
		// METADATA
		CreatedAt: metadata.CreatedAt,
		UpdatedAt: metadata.UpdatedAt,
		DeletedAt: metadata.DeletedAt,
	}
}

func ToWithdrawals(records []*model.Withdrawals) []*withdrawal.Withdrawal {
	entities := make([]*withdrawal.Withdrawal, 0)
	for _, record := range records {
		entities = append(entities, ToWithdrawal(record))
	}
	return entities
}

func FromWithdrawals(entities []withdrawal.Withdrawal) []model.Withdrawals {
	records := make([]model.Withdrawals, 0)
	for _, entity := range entities {
		records = append(records, FromWithdrawal(entity))
	}
	return records
}

func FromWithdrawalDatas(pairs []std.Pair[withdrawal.WithdrawalData, core.Metadata]) []model.Withdrawals {
	records := make([]model.Withdrawals, 0)
	for _, entity := range pairs {
		records = append(records, FromWithdrawalData(entity.First(), entity.Second()))
	}
	return records
}

/* __________________________________________________ */
