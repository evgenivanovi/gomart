package session

import (
	"context"

	"github.com/evgenivanovi/gomart/pkg/search"
)

/* __________________________________________________ */

type SessionKeyRepository interface {
	GetByID(ctx context.Context, id SessionID) (*Session, error)
	FindByID(ctx context.Context, id SessionID) (*Session, error)
	FindByIDs(ctx context.Context, ids []SessionID) ([]*Session, error)
}

type SessionSpecificationRepository interface {
	FindOneBySpec(ctx context.Context, spec search.Specification) (*Session, error)
	FindManyBySpec(ctx context.Context, spec search.Specification) ([]*Session, error)
}

type SessionReadRepository interface {
	SessionKeyRepository
	SessionSpecificationRepository
}

type SessionSaveRepository interface {
	NonAutoSave(ctx context.Context, data Session) (*Session, error)
	NonAutoSaveAll(ctx context.Context, data []Session) error
}

type SessionWriteRepository interface {
	SessionSaveRepository
}

type SessionRepository interface {
	SessionReadRepository
	SessionWriteRepository
}

type SessionRepositoryService struct {
	readRepository  SessionReadRepository
	writeRepository SessionWriteRepository
}

func ProvideSessionRepositoryService(
	readRepository SessionReadRepository,
	writeRepository SessionWriteRepository,
) *SessionRepositoryService {
	return &SessionRepositoryService{
		readRepository:  readRepository,
		writeRepository: writeRepository,
	}
}

func (r *SessionRepositoryService) GetByID(
	ctx context.Context, id SessionID,
) (*Session, error) {
	return r.readRepository.GetByID(ctx, id)
}

func (r *SessionRepositoryService) FindByID(
	ctx context.Context, id SessionID,
) (*Session, error) {
	return r.readRepository.FindByID(ctx, id)
}

func (r *SessionRepositoryService) FindByIDs(
	ctx context.Context, ids []SessionID,
) ([]*Session, error) {
	return r.readRepository.FindByIDs(ctx, ids)
}

func (r *SessionRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*Session, error) {
	return r.readRepository.FindOneBySpec(ctx, spec)
}

func (r *SessionRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*Session, error) {
	return r.readRepository.FindManyBySpec(ctx, spec)
}

func (r *SessionRepositoryService) NonAutoSave(
	ctx context.Context, data Session,
) (*Session, error) {
	return r.writeRepository.NonAutoSave(ctx, data)
}

func (r *SessionRepositoryService) NonAutoSaveAll(
	ctx context.Context, data []Session,
) error {
	return r.writeRepository.NonAutoSaveAll(ctx, data)
}

/* __________________________________________________ */
