package balance

import (
	"context"

	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type BalanceKeyRepository interface {
	GetByID(ctx context.Context, id BalanceID) (*Balance, error)
	FindByID(ctx context.Context, id BalanceID) (*Balance, error)
	FindByIDs(ctx context.Context, ids []BalanceID) ([]*Balance, error)
}

type BalanceSpecificationRepository interface {
	FindOneBySpec(ctx context.Context, spec search.Specification) (*Balance, error)
	FindManyBySpec(ctx context.Context, spec search.Specification) ([]*Balance, error)

	FindOneBySpecExclusively(ctx context.Context, spec search.Specification) (*Balance, error)
	FindManyBySpecExclusively(ctx context.Context, spec search.Specification) ([]*Balance, error)
}

type BalanceReadRepository interface {
	BalanceKeyRepository
	BalanceSpecificationRepository
}

/* __________________________________________________ */

type BalanceSaveRepository interface {
	AutoSave(ctx context.Context, data BalanceData) (*Balance, error)
}

type BalanceUpdateRepository interface {
	Update(ctx context.Context, data Balance) (*Balance, error)
}

type BalanceWriteRepository interface {
	BalanceSaveRepository
	BalanceUpdateRepository
}

/* __________________________________________________ */

type BalanceRepository interface {
	BalanceReadRepository
	BalanceWriteRepository
}

type BalanceRepositoryService struct {
	readRepository  BalanceReadRepository
	writeRepository BalanceWriteRepository
}

func ProvideBalanceRepositoryService(
	readRepository BalanceReadRepository,
	writeRepository BalanceWriteRepository,
) *BalanceRepositoryService {
	return &BalanceRepositoryService{
		readRepository:  readRepository,
		writeRepository: writeRepository,
	}
}

func (r *BalanceRepositoryService) GetByID(
	ctx context.Context, id BalanceID,
) (*Balance, error) {
	return r.readRepository.GetByID(ctx, id)
}

func (r *BalanceRepositoryService) FindByID(
	ctx context.Context, id BalanceID,
) (*Balance, error) {
	return r.readRepository.FindByID(ctx, id)
}

func (r *BalanceRepositoryService) FindByIDs(
	ctx context.Context, ids []BalanceID,
) ([]*Balance, error) {
	return r.readRepository.FindByIDs(ctx, ids)
}

func (r *BalanceRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*Balance, error) {
	return r.readRepository.FindOneBySpec(ctx, spec)
}

func (r *BalanceRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*Balance, error) {
	return r.readRepository.FindManyBySpec(ctx, spec)
}

func (r *BalanceRepositoryService) FindOneBySpecExclusively(
	ctx context.Context, spec search.Specification,
) (*Balance, error) {
	return r.readRepository.FindOneBySpecExclusively(ctx, spec)
}

func (r *BalanceRepositoryService) FindManyBySpecExclusively(
	ctx context.Context, spec search.Specification,
) ([]*Balance, error) {
	return r.readRepository.FindManyBySpecExclusively(ctx, spec)
}

func (r *BalanceRepositoryService) AutoSave(
	ctx context.Context, data BalanceData,
) (*Balance, error) {
	return r.writeRepository.AutoSave(ctx, data)
}

func (r *BalanceRepositoryService) Update(
	ctx context.Context, datas Balance,
) (*Balance, error) {
	return r.writeRepository.Update(ctx, datas)
}

/* __________________________________________________ */
