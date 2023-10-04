package user

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/auth"
	"github.com/evgenivanovi/gomart/internal/domain/auth/session"
	"github.com/evgenivanovi/gomart/internal/domain/auth/token"
	"github.com/evgenivanovi/gomart/internal/domain/balance"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/domain/withdraw"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type UserAuthManager interface {
	Signin(
		ctx context.Context,
		credentials auth.Credentials,
	) (*AuthUser, error)

	Signup(
		ctx context.Context,
		credentials auth.Credentials,
	) (*User, error)

	SignupAndLogin(
		ctx context.Context,
		credentials auth.Credentials,
	) (*AuthUser, error)
}

/* __________________________________________________ */

type UserAuthManagerService struct {
	transactor      common.Transactor
	repository      UserAuthRepository
	passwordManager PasswordManager
	tokenManager    token.TokenManager
	sessionManager  session.SessionManager
	balanceManager  balance.BalanceManager
	withdrawManager withdraw.WithdrawManager
}

func ProvideUserAuthManagerService(
	transactor common.Transactor,
	repository UserAuthRepository,
	passwordManager PasswordManager,
	tokenManager token.TokenManager,
	sessionManager session.SessionManager,
	balanceManager balance.BalanceManager,
	withdrawManager withdraw.WithdrawManager,
) *UserAuthManagerService {
	return &UserAuthManagerService{
		transactor:      transactor,
		repository:      repository,
		passwordManager: passwordManager,
		tokenManager:    tokenManager,
		sessionManager:  sessionManager,
		balanceManager:  balanceManager,
		withdrawManager: withdrawManager,
	}
}

func (u *UserAuthManagerService) Signin(
	ctx context.Context,
	credentials auth.Credentials,
) (*AuthUser, error) {

	user, err := u.getUser(ctx, credentials)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errx.NewErrorWithEntityCode(ErrorUserEntity, core.ErrorNotFoundCode)
	}

	if !u.passwordManager.CompareHashPasswordCtx(
		ctx, credentials.Password(), user.Data().Credentials.Password(),
	) {
		return nil, errx.NewErrorWithEntityCode(ErrorUserEntity, core.ErrorUnauthenticatedCode)
	}

	tokens := u.generateTokens(ctx, *user)

	sess, err := u.createSession(ctx, *user, tokens)
	if err != nil {
		return nil, err
	}

	return NewAuthUser(
		user.Identity(),
		*NewAuthUserData(sess.Identity(), tokens, user.Data()),
		user.Metadata(),
	), nil

}

func (u *UserAuthManagerService) Signup(
	ctx context.Context,
	credentials auth.Credentials,
) (*User, error) {

	user, err := u.createUser(ctx, credentials)
	if err != nil {
		return nil, err
	}

	_, err = u.balanceManager.Create(ctx, user.id)
	if err != nil {
		return nil, err
	}

	_, err = u.withdrawManager.Create(ctx, user.id)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (u *UserAuthManagerService) SignupAndLogin(
	ctx context.Context,
	credentials auth.Credentials,
) (*AuthUser, error) {

	user, err := u.Signup(ctx, credentials)
	if err != nil {
		return nil, err
	}

	tokens := u.generateTokens(ctx, *user)

	sess, err := u.createSession(ctx, *user, tokens)
	if err != nil {
		return nil, err
	}

	return NewAuthUser(
		user.Identity(),
		*NewAuthUserData(sess.Identity(), tokens, user.Data()),
		user.Metadata(),
	), nil

}

/* __________________________________________________ */

func (u *UserAuthManagerService) createUser(
	ctx context.Context, credentials auth.Credentials,
) (*User, error) {

	data := NewUserData(
		*credentials.WithHash(u.passwordHasher()),
	)

	user, err := u.repository.AutoSave(
		ctx, *data, *core.NewNowMetadata(),
	)

	if err != nil {
		return nil, err
	}

	return user, err

}

func (u *UserAuthManagerService) getUser(
	ctx context.Context, credentials auth.Credentials,
) (*User, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(UsernameCondition(credentials.Username()))
	return u.repository.FindOneBySpec(ctx, spec)
}

func (u *UserAuthManagerService) passwordHasher() func(string) string {
	return func(password string) string {
		hash, _ := u.passwordManager.GenerateHashPasswordCtx(context.Background(), password)
		return hash
	}
}

func (u *UserAuthManagerService) generateTokens(
	ctx context.Context, user User,
) token.Tokens {
	tokenData := token.NewTokenData(user.Identity())
	return u.tokenManager.Generate(ctx, *tokenData)
}

func (u *UserAuthManagerService) createSession(
	ctx context.Context, user User, token token.Tokens,
) (session.Session, error) {
	sessionData := session.NewSessionData(user.Identity(), token.RefreshToken)
	return u.sessionManager.Create(ctx, *sessionData)
}

/* __________________________________________________ */
