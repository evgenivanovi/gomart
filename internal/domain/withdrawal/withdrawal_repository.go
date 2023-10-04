package withdrawal

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type WithdrawalKeyRepository interface {
	GetByID(ctx context.Context, id WithdrawalID) (*Withdrawal, error)
	FindByID(ctx context.Context, id WithdrawalID) (*Withdrawal, error)
	FindByIDs(ctx context.Context, ids []WithdrawalID) ([]*Withdrawal, error)
}

type WithdrawalSpecificationRepository interface {
	FindOneBySpec(ctx context.Context, spec search.Specification) (*Withdrawal, error)
	FindManyBySpec(ctx context.Context, spec search.Specification) ([]*Withdrawal, error)

	FindOneBySpecExclusively(ctx context.Context, spec search.Specification) (*Withdrawal, error)
	FindManyBySpecExclusively(ctx context.Context, spec search.Specification) ([]*Withdrawal, error)
}

type WithdrawalReadRepository interface {
	WithdrawalKeyRepository
	WithdrawalSpecificationRepository
}

/* __________________________________________________ */

type WithdrawalSaveRepository interface {
	AutoSave(ctx context.Context, data WithdrawalData, metadata core.Metadata) (*Withdrawal, error)
}

type WithdrawalWriteRepository interface {
	WithdrawalSaveRepository
}

/* __________________________________________________ */

type WithdrawalRepository interface {
	WithdrawalReadRepository
	WithdrawalWriteRepository
}

type WithdrawalRepositoryService struct {
	readRepository  WithdrawalReadRepository
	writeRepository WithdrawalWriteRepository
}

func ProvideWithdrawalRepositoryService(
	readRepository WithdrawalReadRepository,
	writeRepository WithdrawalWriteRepository,
) *WithdrawalRepositoryService {
	return &WithdrawalRepositoryService{
		readRepository:  readRepository,
		writeRepository: writeRepository,
	}
}

func (r *WithdrawalRepositoryService) GetByID(
	ctx context.Context, id WithdrawalID,
) (*Withdrawal, error) {
	return r.readRepository.GetByID(ctx, id)
}

func (r *WithdrawalRepositoryService) FindByID(
	ctx context.Context, id WithdrawalID,
) (*Withdrawal, error) {
	return r.readRepository.FindByID(ctx, id)
}

func (r *WithdrawalRepositoryService) FindByIDs(
	ctx context.Context, ids []WithdrawalID,
) ([]*Withdrawal, error) {
	return r.readRepository.FindByIDs(ctx, ids)
}

func (r *WithdrawalRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*Withdrawal, error) {
	return r.readRepository.FindOneBySpec(ctx, spec)
}

func (r *WithdrawalRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*Withdrawal, error) {
	return r.readRepository.FindManyBySpec(ctx, spec)
}

func (r *WithdrawalRepositoryService) FindOneBySpecExclusively(
	ctx context.Context, spec search.Specification,
) (*Withdrawal, error) {
	return r.readRepository.FindOneBySpecExclusively(ctx, spec)
}

func (r *WithdrawalRepositoryService) FindManyBySpecExclusively(
	ctx context.Context, spec search.Specification,
) ([]*Withdrawal, error) {
	return r.readRepository.FindManyBySpecExclusively(ctx, spec)
}

func (r *WithdrawalRepositoryService) AutoSave(
	ctx context.Context, data WithdrawalData, metadata core.Metadata,
) (*Withdrawal, error) {
	return r.writeRepository.AutoSave(ctx, data, metadata)
}

/* __________________________________________________ */
