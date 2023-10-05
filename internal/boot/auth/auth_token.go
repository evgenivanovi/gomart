package auth

import (
	"github.com/evgenivanovi/gomart/internal/boot/app"

	dm_token "github.com/evgenivanovi/gomart/internal/domain/auth/token"
	authtool "github.com/evgenivanovi/gomart/internal/util/auth"
	"github.com/evgenivanovi/gomart/pkg/fw"
)

/* __________________________________________________ */

var (
	TokenSettings authtool.TokenSettings
	TokenManager  dm_token.TokenManager
)

/* __________________________________________________ */

func BootAuthToken() {

	tokenAccessSecret := app.TokenAccessSecretKeyProperty.
		Calc(fw.FirstStringNotEmpty(authtool.TokenAccessSecretKey))

	tokenAccessTimeExpiration := app.TokenAccessTimeExpirationProperty.
		CalcValue(fw.FirstDurationOr(authtool.TokenAccessExpirationTime)).
		GetDuration()

	tokenRefreshSecret := app.TokenRefreshSecretKeyProperty.
		Calc(fw.FirstStringNotEmpty(authtool.TokenRefreshSecretKey))

	tokenRefreshTimeExpiration := app.TokenRefreshTimeExpirationProperty.
		CalcValue(fw.FirstDurationOr(authtool.TokenRefreshExpirationTime)).
		GetDuration()

	TokenSettings = *authtool.NewTokenSettings(
		authtool.WithAccessSecret(tokenAccessSecret),
		authtool.WithAccessExpiration(tokenAccessTimeExpiration),

		authtool.WithRefreshSecret(tokenRefreshSecret),
		authtool.WithAccessExpiration(tokenRefreshTimeExpiration),
	)

	TokenManager = dm_token.ProvideTokenManagerService(TokenSettings)

}

/* __________________________________________________ */
