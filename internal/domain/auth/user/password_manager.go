package user

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/util/auth"
)

/* __________________________________________________ */

type PasswordManager interface {
	GenerateHashPassword(
		password string,
	) (string, error)

	GenerateHashPasswordCtx(
		ctx context.Context, password string,
	) (string, error)

	CompareHashPassword(
		password string, hash string,
	) bool

	CompareHashPasswordCtx(
		ctx context.Context, password string, hash string,
	) bool
}

type PasswordManagerService struct{}

func ProvidePasswordManagerService() *PasswordManagerService {
	return &PasswordManagerService{}
}

func (p *PasswordManagerService) GenerateHashPassword(
	password string,
) (string, error) {
	return p.GenerateHashPasswordCtx(context.Background(), password)
}

func (p *PasswordManagerService) GenerateHashPasswordCtx(
	ctx context.Context, password string,
) (string, error) {
	return auth.GenerateHashPassword(password)
}

func (p *PasswordManagerService) CompareHashPassword(
	password string, hash string,
) bool {
	return p.CompareHashPasswordCtx(context.Background(), password, hash)
}

func (p *PasswordManagerService) CompareHashPasswordCtx(
	ctx context.Context, password string, hash string,
) bool {
	return auth.CompareHashPassword(password, hash)
}

/* __________________________________________________ */
