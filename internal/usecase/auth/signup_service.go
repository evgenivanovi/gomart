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

type AuthSignUpUsecase interface {
	Execute(context.Context, SignUpRequest) (SignUpResponse, error)
}

type AuthSignUpUsecaseService struct {
	transactor common.Transactor
	manager    user.UserAuthManager
}

func ProvideAuthSignUpUsecaseService(
	transactor common.Transactor,
	manager user.UserAuthManager,
) *AuthSignUpUsecaseService {
	return &AuthSignUpUsecaseService{
		transactor: transactor,
		manager:    manager,
	}
}

func (uc *AuthSignUpUsecaseService) Execute(
	ctx context.Context,
	request SignUpRequest,
) (response SignUpResponse, err error) {

	credentials := auth.NewCredentials(
		request.Payload.Credentials.Username,
		request.Payload.Credentials.Password,
	)

	ctx = uc.transactor.StartEx(ctx)
	usr, err := uc.manager.SignupAndLogin(ctx, *credentials)
	uc.transactor.CloseEx(ctx, err)

	if err != nil {
		return uc.toEmptyResponse(), err
	} else {
		return uc.toResponse(*usr), nil
	}

}

/* __________________________________________________ */

func (uc *AuthSignUpUsecaseService) toEmptyResponse() SignUpResponse {
	return SignUpResponse{}
}

func (uc *AuthSignUpUsecaseService) toResponse(user user.AuthUser) SignUpResponse {
	return SignUpResponse{
		Payload: SignUpResponsePayload{
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
