package token

import (
	"time"

	"github.com/evgenivanovi/gomart/internal/domain/common"
)

/* __________________________________________________ */

type AccessToken struct {
	Token     string
	ExpiresAt time.Time
}

func NewAccessToken(token string, expiration time.Time) *AccessToken {
	return &AccessToken{
		Token:     token,
		ExpiresAt: expiration,
	}
}

type RefreshToken struct {
	Token     string
	ExpiresAt time.Time
}

func NewRefreshToken(token string, expiration time.Time) *RefreshToken {
	return &RefreshToken{
		Token:     token,
		ExpiresAt: expiration,
	}
}

/* __________________________________________________ */

type Tokens struct {
	AccessToken  AccessToken
	RefreshToken RefreshToken
}

func NewTokens(access AccessToken, refresh RefreshToken) *Tokens {
	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

/* __________________________________________________ */

type TokenData struct {
	UserID common.UserID
}

func NewTokenData(userID common.UserID) *TokenData {
	return &TokenData{
		UserID: userID,
	}
}

/* __________________________________________________ */
