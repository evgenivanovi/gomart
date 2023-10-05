package auth

import "time"

/* __________________________________________________ */

type AccessToken struct {
	Token     string
	ExpiresAt time.Time
}

type RefreshToken struct {
	Token     string
	ExpiresAt time.Time
}

type Tokens struct {
	AccessToken  AccessToken
	RefreshToken RefreshToken
}

type Session struct {
	SessionID string
	Tokens    Tokens
}

/* __________________________________________________ */
