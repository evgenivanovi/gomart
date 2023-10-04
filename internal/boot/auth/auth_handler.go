package auth

import (
	"github.com/evgenivanovi/gomart/internal/http/infra/auth"
)

/* __________________________________________________ */

var (
	SignInController *auth.SignInController
	SignUpController *auth.SignUpController
)

/* __________________________________________________ */

func BootAuthHandlers() {

	SignInController = auth.ProvideSignInController(
		AuthSignInUsecase,
	)

	SignUpController = auth.ProvideSignUpController(
		AuthSignUpUsecase,
	)

}

/* __________________________________________________ */
