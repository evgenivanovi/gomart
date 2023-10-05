package auth

import (
	"net/http"

	authtool "github.com/evgenivanovi/gomart/internal/util/auth"
	"github.com/evgenivanovi/gomart/pkg/stdx/jwtx"
)

/* __________________________________________________ */

var (
	AuthMW func(next http.Handler) http.Handler
)

/* __________________________________________________ */

func BootAuthMW() {

	AuthMW = jwtx.New(
		jwtx.WithKey(authtool.KeyProvider(TokenSettings.AccessSecret())),
		jwtx.WithMethod(authtool.MethodProvider()),
		jwtx.WithClaims(authtool.ClaimsProvider()),
		jwtx.WithExtractor(authtool.ExtractorProvider()),
		jwtx.WithAfter(authtool.AfterProvider()),
		jwtx.WithRecover(authtool.RecoveryProvider()),
	)

}

/* __________________________________________________ */
