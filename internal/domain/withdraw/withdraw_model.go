package withdraw

import (
	"github.com/evgenivanovi/gomart/internal/domain/common"
)

/* __________________________________________________ */

type WithdrawID struct {
	id int64
}

func NewWithdrawID(id int64) *WithdrawID {
	return &WithdrawID{
		id: id,
	}
}

func (u WithdrawID) ID() int64 {
	return u.id
}

/* __________________________________________________ */

type WithdrawData struct {
	UserID   common.UserID
	Withdraw common.Points
}

func NewWithdrawData(userID common.UserID, Withdraw common.Points) *WithdrawData {
	return &WithdrawData{
		UserID:   userID,
		Withdraw: Withdraw,
	}
}

func NewEmptyWithdrawData(userID common.UserID) *WithdrawData {
	return NewWithdrawData(userID, *common.NewEmptyPoints())
}

/* __________________________________________________ */

type Withdraw struct {
	id   WithdrawID
	data WithdrawData
}

func NewWithdraw(id WithdrawID, data WithdrawData) *Withdraw {
	return &Withdraw{
		id:   id,
		data: data,
	}
}

func NewEmptyWithdraw(id WithdrawID, userID common.UserID) *Withdraw {
	return &Withdraw{
		id:   id,
		data: *NewEmptyWithdrawData(userID),
	}
}

func (e *Withdraw) Identity() WithdrawID {
	return e.id
}

func (e *Withdraw) Data() WithdrawData {
	return e.data
}

func (e *Withdraw) WithIdentity(id WithdrawID) Withdraw {
	return *NewWithdraw(id, e.data)
}

func (e *Withdraw) WithIdentityFn(id func() WithdrawID) Withdraw {
	return *NewWithdraw(id(), e.data)
}

func (e *Withdraw) WithData(data WithdrawData) Withdraw {
	return *NewWithdraw(e.id, data)
}

func (e *Withdraw) WithDataFn(data func() WithdrawData) Withdraw {
	return *NewWithdraw(e.id, data())
}

/* __________________________________________________ */
