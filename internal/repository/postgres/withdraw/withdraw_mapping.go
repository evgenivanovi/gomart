package withdraw

import (
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/withdraw"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
)

/* __________________________________________________ */

func ToWithdraw(record *model.Withdraws) *withdraw.Withdraw {

	id := withdraw.NewWithdrawID(record.ID)

	data := withdraw.NewWithdrawData(
		*common.NewUserID(record.UserID),
		*common.NewPoints(record.Amount),
	)

	return withdraw.NewWithdraw(*id, *data)

}

func FromWithdraw(entity withdraw.Withdraw) model.Withdraws {
	return model.Withdraws{
		ID:     entity.Identity().ID(),
		UserID: entity.Data().UserID.ID(),
		Amount: entity.Data().Withdraw.Amount(),
	}
}

func FromWithdrawData(entity withdraw.WithdrawData) model.Withdraws {
	return model.Withdraws{
		UserID: entity.UserID.ID(),
		Amount: entity.Withdraw.Amount(),
	}
}

func ToWithdraws(records []*model.Withdraws) []*withdraw.Withdraw {
	entities := make([]*withdraw.Withdraw, 0)
	for _, record := range records {
		entities = append(entities, ToWithdraw(record))
	}
	return entities
}

func FromWithdraws(entities []withdraw.Withdraw) []model.Withdraws {
	records := make([]model.Withdraws, 0)
	for _, entity := range entities {
		records = append(records, FromWithdraw(entity))
	}
	return records
}

func FromWithdrawDatas(entities []withdraw.WithdrawData) []model.Withdraws {
	records := make([]model.Withdraws, 0)
	for _, entity := range entities {
		records = append(records, FromWithdrawData(entity))
	}
	return records
}

/* __________________________________________________ */
