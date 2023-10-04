package withdrawal

import (
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
)

/* __________________________________________________ */

type WithdrawalID struct {
	id int64
}

func NewWithdrawalID(id int64) *WithdrawalID {
	return &WithdrawalID{
		id: id,
	}
}

func (u WithdrawalID) ID() int64 {
	return u.id
}

/* __________________________________________________ */

type WithdrawalData struct {
	UserID common.UserID
	Amount common.Points
	Order  string
}

func NewWithdrawalData(userID common.UserID, amount common.Points, order string) *WithdrawalData {
	return &WithdrawalData{
		UserID: userID,
		Amount: amount,
		Order:  order,
	}
}

/* __________________________________________________ */

type Withdrawal struct {
	id       WithdrawalID
	data     WithdrawalData
	metadata core.Metadata
}

func NewWithdrawal(id WithdrawalID, data WithdrawalData, metadata core.Metadata) *Withdrawal {
	return &Withdrawal{
		id:       id,
		data:     data,
		metadata: metadata,
	}
}

func ToWithdrawalPointers(entities []Withdrawal) []*Withdrawal {
	result := make([]*Withdrawal, 0)
	for _, entity := range entities {
		result = append(result, &entity)
	}
	return result
}

func ToWithdrawalValues(entities []*Withdrawal) []Withdrawal {
	result := make([]Withdrawal, 0)
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

func (e *Withdrawal) Identity() WithdrawalID {
	return e.id
}

func (e *Withdrawal) Data() WithdrawalData {
	return e.data
}

func (e *Withdrawal) Metadata() core.Metadata {
	return e.metadata
}

/* __________________________________________________ */

type WithdrawalHistory struct {
	UserID      common.UserID
	Withdrawals []Withdrawal
}

func NewWithdrawalHistory(
	userID common.UserID,
	withdrawals []Withdrawal,
) *WithdrawalHistory {
	return &WithdrawalHistory{
		UserID:      userID,
		Withdrawals: withdrawals,
	}
}

/* __________________________________________________ */
