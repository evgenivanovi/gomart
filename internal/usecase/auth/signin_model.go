package auth

import (
	"github.com/evgenivanovi/gomart/internal/model/auth"
)

/* __________________________________________________ */

type SignInRequestPayload struct {
	Credentials auth.Credentials
}

type SignInRequest struct {
	Payload SignInRequestPayload
}

/* __________________________________________________ */

type SignInResponsePayload struct {
	Session auth.Session
	User    auth.User
}

type SignInResponse struct {
	Payload SignInResponsePayload
}

/* __________________________________________________ */
