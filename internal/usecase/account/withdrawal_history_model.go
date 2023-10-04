package account

import "github.com/evgenivanovi/gomart/internal/model/withdraw"

/* __________________________________________________ */

type WithdrawalHistoryRequestPayload struct {
	UserID int64
}

type WithdrawalHistoryRequest struct {
	Payload WithdrawalHistoryRequestPayload
}

/* __________________________________________________ */

type WithdrawalHistoryResponsePayload struct {
	Withdrawals withdraw.WithdrawalHistory
}

type WithdrawalHistoryResponse struct {
	Payload WithdrawalHistoryResponsePayload
}

/* __________________________________________________ */
