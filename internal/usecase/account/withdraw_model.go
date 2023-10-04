package account

import "github.com/evgenivanovi/gomart/internal/model/withdraw"

/* __________________________________________________ */

type WithdrawRequestPayload struct {
	UserID int64
	Order  string
	Amount float64
}

type WithdrawRequest struct {
	Payload WithdrawRequestPayload
}

/* __________________________________________________ */

type WithdrawResponsePayload struct {
	Withdrawal withdraw.UserWithdrawal
}

type WithdrawResponse struct {
	Payload WithdrawResponsePayload
}

/* __________________________________________________ */
