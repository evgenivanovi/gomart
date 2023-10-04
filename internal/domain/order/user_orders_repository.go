package order

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type UserOrdersKeyRepository interface {
	GetByID(ctx context.Context, id common.UserID) (*UserOrders, error)
	FindByID(ctx context.Context, id common.UserID) (*UserOrders, error)
}

type UserOrdersSpecificationRepository interface {
	FindBySpec(ctx context.Context, spec search.Specification) (*UserOrders, error)
}

type UserOrdersReadRepository interface {
	UserOrdersKeyRepository
	UserOrdersSpecificationRepository
}

/* __________________________________________________ */

type UserOrdersRepository interface {
	UserOrdersReadRepository
}

type UserOrdersRepositoryService struct {
	readRepository UserOrdersReadRepository
}

func ProvideUserOrdersRepositoryService(
	readRepository UserOrdersReadRepository,
) *UserOrdersRepositoryService {
	return &UserOrdersRepositoryService{
		readRepository: readRepository,
	}
}

func (r *UserOrdersRepositoryService) GetByID(
	ctx context.Context, id common.UserID,
) (*UserOrders, error) {
	return r.readRepository.GetByID(ctx, id)
}

func (r *UserOrdersRepositoryService) FindByID(
	ctx context.Context, id common.UserID,
) (*UserOrders, error) {
	return r.readRepository.FindByID(ctx, id)
}

func (r *UserOrdersRepositoryService) FindBySpec(
	ctx context.Context, spec search.Specification,
) (*UserOrders, error) {
	return r.readRepository.FindBySpec(ctx, spec)
}

/* __________________________________________________ */
