package balance

import (
	"github.com/evgenivanovi/gomart/internal/domain/balance"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
)

/* __________________________________________________ */

func ToBalance(record *model.Balances) *balance.Balance {

	id := balance.NewBalanceID(record.ID)

	data := balance.NewBalanceData(
		*common.NewUserID(record.UserID),
		*common.NewPoints(record.Amount),
	)

	return balance.NewBalance(*id, *data)

}

func FromBalance(entity balance.Balance) model.Balances {
	return model.Balances{
		ID:     entity.Identity().ID(),
		UserID: entity.Data().UserID.ID(),
		Amount: entity.Data().Balance.Amount(),
	}
}

func FromBalanceData(entity balance.BalanceData) model.Balances {
	return model.Balances{
		UserID: entity.UserID.ID(),
		Amount: entity.Balance.Amount(),
	}
}

func ToBalances(records []*model.Balances) []*balance.Balance {
	entities := make([]*balance.Balance, 0)
	for _, record := range records {
		entities = append(entities, ToBalance(record))
	}
	return entities
}

func FromBalances(entities []balance.Balance) []model.Balances {
	records := make([]model.Balances, 0)
	for _, entity := range entities {
		records = append(records, FromBalance(entity))
	}
	return records
}

func FromBalanceDatas(entities []balance.BalanceData) []model.Balances {
	records := make([]model.Balances, 0)
	for _, entity := range entities {
		records = append(records, FromBalanceData(entity))
	}
	return records
}

/* __________________________________________________ */
