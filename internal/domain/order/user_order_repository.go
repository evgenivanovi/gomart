package order

import (
	"context"

	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type UserOrderKeyRepository interface {
	GetByID(ctx context.Context, id OrderID) (*UserOrder, error)
	FindByID(ctx context.Context, id OrderID) (*UserOrder, error)
	FindByIDs(ctx context.Context, ids []OrderID) ([]*UserOrder, error)
}

type UserOrderSpecificationRepository interface {
	FindOneBySpec(ctx context.Context, spec search.Specification) (*UserOrder, error)
	FindManyBySpec(ctx context.Context, spec search.Specification) ([]*UserOrder, error)

	FindOneBySpecExclusively(ctx context.Context, spec search.Specification) (*UserOrder, error)
	FindManyBySpecExclusively(ctx context.Context, spec search.Specification) ([]*UserOrder, error)
}

type UserOrderReadRepository interface {
	UserOrderKeyRepository
	UserOrderSpecificationRepository
}

/* __________________________________________________ */

type UserOrderSaveRepository interface {
	NonAutoSave(ctx context.Context, data UserOrder) (*UserOrder, error)
	NonAutoSaveAll(ctx context.Context, datas []UserOrder) error
}

type UserOrderUpdateRepository interface {
	Update(ctx context.Context, data UserOrder) (*UserOrder, error)
}

type UserOrderWriteRepository interface {
	UserOrderSaveRepository
	UserOrderUpdateRepository
}

/* __________________________________________________ */

type UserOrderRepository interface {
	UserOrderReadRepository
	UserOrderWriteRepository
}

type UserOrderRepositoryService struct {
	readRepository  UserOrderReadRepository
	writeRepository UserOrderWriteRepository
}

func ProvideUserOrderRepositoryService(
	readRepository UserOrderReadRepository,
	writeRepository UserOrderWriteRepository,
) *UserOrderRepositoryService {
	return &UserOrderRepositoryService{
		readRepository:  readRepository,
		writeRepository: writeRepository,
	}
}

func (r *UserOrderRepositoryService) GetByID(
	ctx context.Context, id OrderID,
) (*UserOrder, error) {
	return r.readRepository.GetByID(ctx, id)
}

func (r *UserOrderRepositoryService) FindByID(
	ctx context.Context, id OrderID,
) (*UserOrder, error) {
	return r.readRepository.FindByID(ctx, id)
}

func (r *UserOrderRepositoryService) FindByIDs(
	ctx context.Context, ids []OrderID,
) ([]*UserOrder, error) {
	return r.readRepository.FindByIDs(ctx, ids)
}

func (r *UserOrderRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*UserOrder, error) {
	return r.readRepository.FindOneBySpec(ctx, spec)
}

func (r *UserOrderRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*UserOrder, error) {
	return r.readRepository.FindManyBySpec(ctx, spec)
}

func (r *UserOrderRepositoryService) FindOneBySpecExclusively(
	ctx context.Context, spec search.Specification,
) (*UserOrder, error) {
	return r.readRepository.FindOneBySpecExclusively(ctx, spec)
}

func (r *UserOrderRepositoryService) FindManyBySpecExclusively(
	ctx context.Context, spec search.Specification,
) ([]*UserOrder, error) {
	return r.readRepository.FindManyBySpecExclusively(ctx, spec)
}

func (r *UserOrderRepositoryService) NonAutoSave(
	ctx context.Context, data UserOrder,
) (*UserOrder, error) {
	return r.writeRepository.NonAutoSave(ctx, data)
}

func (r *UserOrderRepositoryService) NonAutoSaveAll(
	ctx context.Context, datas []UserOrder,
) error {
	return r.writeRepository.NonAutoSaveAll(ctx, datas)
}

func (r *UserOrderRepositoryService) Update(
	ctx context.Context, datas UserOrder,
) (*UserOrder, error) {
	return r.writeRepository.Update(ctx, datas)
}

/* __________________________________________________ */
