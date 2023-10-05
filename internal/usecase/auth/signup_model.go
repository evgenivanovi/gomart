package auth

import "github.com/evgenivanovi/gomart/internal/model/auth"

/* __________________________________________________ */

type SignUpRequestPayload struct {
	Credentials auth.Credentials
}

type SignUpRequest struct {
	Payload SignUpRequestPayload
}

/* __________________________________________________ */

type SignUpResponsePayload struct {
	Session auth.Session
	User    auth.User
}

type SignUpResponse struct {
	Payload SignUpResponsePayload
}

/* __________________________________________________ */
