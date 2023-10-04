package account

/* __________________________________________________ */

type GetBalanceRequestPayload struct {
	UserID int64
}

type GetBalanceRequest struct {
	Payload GetBalanceRequestPayload
}

/* __________________________________________________ */

type GetBalanceResponsePayload struct {
	Balance   float64
	Withdrawn float64
}

type GetBalanceResponse struct {
	Payload GetBalanceResponsePayload
}

/* __________________________________________________ */
