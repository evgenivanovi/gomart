package withdraw

import md "github.com/evgenivanovi/gomart/internal/model"

/* __________________________________________________ */

type Withdrawal struct {
	Order    string
	Amount   float64
	Metadata md.Metadata
}

type UserWithdrawal struct {
	UserID   int64
	Order    string
	Amount   int64
	Metadata md.Metadata
}

type WithdrawalHistory struct {
	UserID      int64
	Withdrawals []Withdrawal
}

/* __________________________________________________ */
