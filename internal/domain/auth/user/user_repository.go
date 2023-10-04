package user

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/pkg/search"
	"github.com/evgenivanovi/gomart/pkg/std"
)

/* __________________________________________________ */

type UserAuthKeyRepository interface {
	GetByID(ctx context.Context, id common.UserID) (*User, error)
	FindByID(ctx context.Context, id common.UserID) (*User, error)
	FindByIDs(ctx context.Context, ids []common.UserID) ([]*User, error)
}

type UserAuthSpecificationRepository interface {
	FindOneBySpec(ctx context.Context, spec search.Specification) (*User, error)
	FindManyBySpec(ctx context.Context, spec search.Specification) ([]*User, error)
}

type UserAuthReadRepository interface {
	UserAuthKeyRepository
	UserAuthSpecificationRepository
}

/* __________________________________________________ */

type UserAuthSaveRepository interface {
	AutoSave(ctx context.Context, data UserData, metadata core.Metadata) (*User, error)
	AutoSaveAll(ctx context.Context, datas []std.Pair[UserData, core.Metadata]) error
}

type UserAuthWriteRepository interface {
	UserAuthSaveRepository
}

/* __________________________________________________ */

type UserAuthRepository interface {
	UserAuthReadRepository
	UserAuthWriteRepository
}

type UserAuthRepositoryService struct {
	readRepository  UserAuthReadRepository
	writeRepository UserAuthWriteRepository
}

func ProvideUserAuthRepositoryService(
	readRepository UserAuthReadRepository,
	writeRepository UserAuthWriteRepository,
) *UserAuthRepositoryService {
	return &UserAuthRepositoryService{
		readRepository:  readRepository,
		writeRepository: writeRepository,
	}
}

func (r *UserAuthRepositoryService) GetByID(
	ctx context.Context, id common.UserID,
) (*User, error) {
	return r.readRepository.GetByID(ctx, id)
}

func (r *UserAuthRepositoryService) FindByID(
	ctx context.Context, id common.UserID,
) (*User, error) {
	return r.readRepository.FindByID(ctx, id)
}

func (r *UserAuthRepositoryService) FindByIDs(
	ctx context.Context, ids []common.UserID,
) ([]*User, error) {
	return r.readRepository.FindByIDs(ctx, ids)
}

func (r *UserAuthRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*User, error) {
	return r.readRepository.FindOneBySpec(ctx, spec)
}

func (r *UserAuthRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*User, error) {
	return r.readRepository.FindManyBySpec(ctx, spec)
}

func (r *UserAuthRepositoryService) AutoSave(
	ctx context.Context, data UserData, metadata core.Metadata,
) (*User, error) {
	return r.writeRepository.AutoSave(ctx, data, metadata)
}

func (r *UserAuthRepositoryService) AutoSaveAll(
	ctx context.Context, datas []std.Pair[UserData, core.Metadata],
) error {
	return r.writeRepository.AutoSaveAll(ctx, datas)
}

/* __________________________________________________ */
