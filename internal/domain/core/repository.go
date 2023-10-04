package core

import (
	"context"

	"github.com/evgenivanovi/gomart/pkg/search"
	"github.com/evgenivanovi/gomart/pkg/std"
)

/* __________________________________________________ */

//goland:noinspection GoSnakeCaseUsage
type KeyRepository[ID any, ENTITY Entity[ID, any]] interface {
	FindByID(ctx context.Context, id ID) (*ENTITY, error)
	FindByIDs(ctx context.Context, ids []ID) (*[]ENTITY, error)
}

type SnapshotRepository[ID any, ENTITY Entity[ID, any]] interface {
	GetAll(ctx context.Context) ([]*ENTITY, error)
}

//goland:noinspection GoSnakeCaseUsage
type AutoSaveRepository[ID any, ENTITY Entity[ID, ENTITY_DATA], ENTITY_DATA any] interface {
	AutoSave(ctx context.Context, data ENTITY_DATA) (*ENTITY, error)
	AutoSaveAll(ctx context.Context, datas []std.Pair[ENTITY_DATA, Metadata]) error
}

//goland:noinspection GoSnakeCaseUsage
type NonAutoSaveRepository[ID any, ENTITY Entity[ID, any]] interface {
	NonAutoSave(ctx context.Context, data ENTITY) (*ENTITY, error)
	NonAutoSaveAll(ctx context.Context, data []*ENTITY) error
}

//goland:noinspection GoSnakeCaseUsage
type DeleteRepository[ID any, ENTITY Entity[ID, any]] interface {
	Remove(ctx context.Context, id ID) (*ENTITY, error)
	RemoveAll(ctx context.Context, ids []ID) error
	RemoveBySpec(ctx context.Context, specification search.Specification) error
}

//goland:noinspection GoSnakeCaseUsage
type SpecificationRepository[ID any, ENTITY Entity[ID, any]] interface {
	FindOneBySpec(ctx context.Context, specification search.Specification) (*ENTITY, error)
	FindManyBySpec(ctx context.Context, specification search.Specification) ([]*ENTITY, error)
	CountBySpec(ctx context.Context, specification search.Specification) (int64, error)
}

/* __________________________________________________ */
