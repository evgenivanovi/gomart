package withdraw

import (
	"context"

	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type WithdrawKeyRepository interface {
	GetByID(ctx context.Context, id WithdrawID) (*Withdraw, error)
	FindByID(ctx context.Context, id WithdrawID) (*Withdraw, error)
	FindByIDs(ctx context.Context, ids []WithdrawID) ([]*Withdraw, error)
}

type WithdrawSpecificationRepository interface {
	FindOneBySpec(ctx context.Context, spec search.Specification) (*Withdraw, error)
	FindManyBySpec(ctx context.Context, spec search.Specification) ([]*Withdraw, error)

	FindOneBySpecExclusively(ctx context.Context, spec search.Specification) (*Withdraw, error)
	FindManyBySpecExclusively(ctx context.Context, spec search.Specification) ([]*Withdraw, error)
}

type WithdrawReadRepository interface {
	WithdrawKeyRepository
	WithdrawSpecificationRepository
}

/* __________________________________________________ */

type WithdrawSaveRepository interface {
	AutoSave(ctx context.Context, data WithdrawData) (*Withdraw, error)
}

type WithdrawUpdateRepository interface {
	Update(ctx context.Context, data Withdraw) (*Withdraw, error)
}

type WithdrawWriteRepository interface {
	WithdrawSaveRepository
	WithdrawUpdateRepository
}

/* __________________________________________________ */

type WithdrawRepository interface {
	WithdrawReadRepository
	WithdrawWriteRepository
}

type WithdrawRepositoryService struct {
	readRepository  WithdrawReadRepository
	writeRepository WithdrawWriteRepository
}

func ProvideWithdrawRepositoryService(
	readRepository WithdrawReadRepository,
	writeRepository WithdrawWriteRepository,
) *WithdrawRepositoryService {
	return &WithdrawRepositoryService{
		readRepository:  readRepository,
		writeRepository: writeRepository,
	}
}

func (r *WithdrawRepositoryService) GetByID(
	ctx context.Context, id WithdrawID,
) (*Withdraw, error) {
	return r.readRepository.GetByID(ctx, id)
}

func (r *WithdrawRepositoryService) FindByID(
	ctx context.Context, id WithdrawID,
) (*Withdraw, error) {
	return r.readRepository.FindByID(ctx, id)
}

func (r *WithdrawRepositoryService) FindByIDs(
	ctx context.Context, ids []WithdrawID,
) ([]*Withdraw, error) {
	return r.readRepository.FindByIDs(ctx, ids)
}

func (r *WithdrawRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*Withdraw, error) {
	return r.readRepository.FindOneBySpec(ctx, spec)
}

func (r *WithdrawRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*Withdraw, error) {
	return r.readRepository.FindManyBySpec(ctx, spec)
}

func (r *WithdrawRepositoryService) FindOneBySpecExclusively(
	ctx context.Context, spec search.Specification,
) (*Withdraw, error) {
	return r.readRepository.FindOneBySpecExclusively(ctx, spec)
}

func (r *WithdrawRepositoryService) FindManyBySpecExclusively(
	ctx context.Context, spec search.Specification,
) ([]*Withdraw, error) {
	return r.readRepository.FindManyBySpecExclusively(ctx, spec)
}

func (r *WithdrawRepositoryService) AutoSave(
	ctx context.Context, data WithdrawData,
) (*Withdraw, error) {
	return r.writeRepository.AutoSave(ctx, data)
}

func (r *WithdrawRepositoryService) Update(
	ctx context.Context, datas Withdraw,
) (*Withdraw, error) {
	return r.writeRepository.Update(ctx, datas)
}

/* __________________________________________________ */
