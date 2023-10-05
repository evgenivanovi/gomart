package token

import (
	"context"

	authtool "github.com/evgenivanovi/gomart/internal/util/auth"
	"github.com/evgenivanovi/gomart/pkg/std/time"
	"github.com/evgenivanovi/gomart/pkg/stdx/jwtx"
)

/* __________________________________________________ */

type TokenManager interface {
	Generate(ctx context.Context, data TokenData) Tokens
}

type TokenManagerService struct {
	settings authtool.TokenSettings
}

func ProvideTokenManagerService(
	settings authtool.TokenSettings,
) *TokenManagerService {
	return &TokenManagerService{
		settings: settings,
	}
}

func (t *TokenManagerService) Generate(
	ctx context.Context, data TokenData,
) Tokens {

	accessToken := CreateAccessToken(data, t.settings.AccessExpiration())
	accessTokenString, _ := jwtx.SignJWT(*accessToken, t.settings.AccessSecret())
	access := NewAccessToken(accessTokenString, time.NowPlus(t.settings.AccessExpiration()))

	refreshToken := CreateRefreshToken(data, t.settings.RefreshExpiration())
	refreshTokenString, _ := jwtx.SignJWT(*refreshToken, t.settings.AccessSecret())
	refresh := NewRefreshToken(refreshTokenString, time.NowPlus(t.settings.AccessExpiration()))

	return *NewTokens(*access, *refresh)

}

/* __________________________________________________ */
