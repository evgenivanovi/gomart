package auth

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/auth"
	"github.com/evgenivanovi/gomart/internal/domain/auth/user"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	md "github.com/evgenivanovi/gomart/internal/model"
	auth_md "github.com/evgenivanovi/gomart/internal/model/auth"
)

/* __________________________________________________ */

type AuthSignInUsecase interface {
	Execute(context.Context, SignInRequest) (SignInResponse, error)
}

type AuthSignInUsecaseService struct {
	transactor common.Transactor
	manager    user.UserAuthManager
}

func ProvideAuthSignInUsecaseService(
	transactor common.Transactor,
	manager user.UserAuthManager,
) *AuthSignInUsecaseService {
	return &AuthSignInUsecaseService{
		transactor: transactor,
		manager:    manager,
	}
}

func (uc *AuthSignInUsecaseService) Execute(
	ctx context.Context,
	request SignInRequest,
) (response SignInResponse, err error) {

	credentials := auth.NewCredentials(
		request.Payload.Credentials.Username,
		request.Payload.Credentials.Password,
	)

	ctx = uc.transactor.StartEx(ctx)
	usr, err := uc.manager.Signin(ctx, *credentials)
	uc.transactor.CloseEx(ctx, err)

	if err != nil {
		return uc.toEmptyResponse(), err
	} else {
		return uc.toResponse(*usr), nil
	}

}

/* __________________________________________________ */

func (uc *AuthSignInUsecaseService) toEmptyResponse() SignInResponse {
	return SignInResponse{}
}

func (uc *AuthSignInUsecaseService) toResponse(user user.AuthUser) SignInResponse {
	return SignInResponse{
		Payload: SignInResponsePayload{
			Session: auth_md.Session{
				SessionID: user.Data().SessionID.ID(),
				Tokens: auth_md.Tokens{
					AccessToken: auth_md.AccessToken{
						Token:     user.Data().AccessToken.Token,
						ExpiresAt: user.Data().AccessToken.ExpiresAt,
					},
					RefreshToken: auth_md.RefreshToken{
						Token:     user.Data().RefreshToken.Token,
						ExpiresAt: user.Data().RefreshToken.ExpiresAt,
					},
				},
			},
			User: auth_md.User{
				ID: user.Identity().ID(),
				Data: auth_md.UserData{
					Credentials: auth_md.Credentials{
						Username: user.Data().Credentials.Username(),
						Password: user.Data().Credentials.Password(),
					},
				},
				Metadata: md.Metadata{
					CreatedAt: user.Metadata().CreatedAt,
					UpdatedAt: user.Metadata().UpdatedAt,
					DeletedAt: user.Metadata().DeletedAt,
				},
			},
		},
	}
}

/* __________________________________________________ */
