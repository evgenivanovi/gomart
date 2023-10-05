package auth

import (
	"github.com/evgenivanovi/gomart/internal/boot/account"
	"github.com/evgenivanovi/gomart/internal/boot/common"
	"github.com/evgenivanovi/gomart/internal/boot/pg"
	user_dm "github.com/evgenivanovi/gomart/internal/domain/auth/user"
	user_pg "github.com/evgenivanovi/gomart/internal/repository/postgres/user"
	auth_uc "github.com/evgenivanovi/gomart/internal/usecase/auth"
)

/* __________________________________________________ */

var (
	UserAuthReadRepository  user_dm.UserAuthReadRepository
	UserAuthWriteRepository user_dm.UserAuthWriteRepository
	UserAuthRepository      user_dm.UserAuthRepository

	PasswordManager user_dm.PasswordManager
	UserAuthManager user_dm.UserAuthManager

	AuthSignInUsecase auth_uc.AuthSignInUsecase
	AuthSignUpUsecase auth_uc.AuthSignUpUsecase
)

/* __________________________________________________ */

func BootAuthUser() {

	UserAuthReadRepository = user_pg.ProvideUserAuthReadRepositoryService(
		pg.PostgresReadRequester,
	)
	UserAuthWriteRepository = user_pg.ProvideUserAuthWriteRepositoryService(
		pg.PostgresWriteRequester,
	)

	PasswordManager = user_dm.ProvidePasswordManagerService()

	UserAuthRepository = user_dm.ProvideUserAuthRepositoryService(
		UserAuthReadRepository, UserAuthWriteRepository,
	)

	UserAuthManager = user_dm.ProvideUserAuthManagerService(
		common.Transactor,
		UserAuthRepository,
		PasswordManager,
		TokenManager,
		SessionManager,
		account.BalanceManager,
		account.WithdrawManager,
	)

	AuthSignInUsecase = auth_uc.ProvideAuthSignInUsecaseService(
		common.Transactor, UserAuthManager,
	)
	AuthSignUpUsecase = auth_uc.ProvideAuthSignUpUsecaseService(
		common.Transactor, UserAuthManager,
	)

}

/* __________________________________________________ */
