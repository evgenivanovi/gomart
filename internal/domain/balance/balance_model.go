package balance

import (
	"github.com/evgenivanovi/gomart/internal/domain/common"
)

/* __________________________________________________ */

type BalanceID struct {
	id int64
}

func NewBalanceID(id int64) *BalanceID {
	return &BalanceID{
		id: id,
	}
}

func (b BalanceID) ID() int64 {
	return b.id
}

/* __________________________________________________ */

type BalanceData struct {
	UserID  common.UserID
	Balance common.Points
}

func NewBalanceData(userID common.UserID, balance common.Points) *BalanceData {
	return &BalanceData{
		UserID:  userID,
		Balance: balance,
	}
}

func NewEmptyBalanceData(userID common.UserID) *BalanceData {
	return NewBalanceData(userID, *common.NewEmptyPoints())
}

/* __________________________________________________ */

type Balance struct {
	id   BalanceID
	data BalanceData
}

func NewBalance(id BalanceID, data BalanceData) *Balance {
	return &Balance{
		id:   id,
		data: data,
	}
}

func NewEmptyBalance(id BalanceID, userID common.UserID) *Balance {
	return &Balance{
		id:   id,
		data: *NewEmptyBalanceData(userID),
	}
}

func (e *Balance) Identity() BalanceID {
	return e.id
}

func (e *Balance) Data() BalanceData {
	return e.data
}

func (e *Balance) WithIdentity(id BalanceID) Balance {
	return *NewBalance(id, e.data)
}

func (e *Balance) WithIdentityFn(id func() BalanceID) Balance {
	return *NewBalance(id(), e.data)
}

func (e *Balance) WithData(data BalanceData) Balance {
	return *NewBalance(e.id, data)
}

func (e *Balance) WithDataFn(data func() BalanceData) Balance {
	return *NewBalance(e.id, data())
}

/* __________________________________________________ */
