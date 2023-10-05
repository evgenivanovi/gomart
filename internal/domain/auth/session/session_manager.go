package session

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/pkg/search"
	"github.com/google/uuid"
)

/* __________________________________________________ */

type SessionIDGenerator interface {
	Generate() SessionID
}

type SessionIDGeneratorService struct{}

func ProvideSessionIDGeneratorService() *SessionIDGeneratorService {
	return &SessionIDGeneratorService{}
}

func (s *SessionIDGeneratorService) Generate() SessionID {
	return *NewSessionID(uuid.NewString())
}

/* __________________________________________________ */

type SessionManager interface {
	Get(ctx context.Context, user common.UserID) (Session, error)
	Create(ctx context.Context, data SessionData) (Session, error)
}

type SessionManagerService struct {
	sequence   SessionIDGenerator
	repository SessionRepository
}

func ProvideSessionManagerService(
	sequence SessionIDGenerator,
	repository SessionRepository,
) *SessionManagerService {
	return &SessionManagerService{
		sequence:   sequence,
		repository: repository,
	}
}

func (s *SessionManagerService) Get(
	ctx context.Context,
	user common.UserID,
) (Session, error) {
	session, err := s.getSession(ctx, user)
	if err != nil {
		return NewEmptySession(), err
	}
	return *session, nil
}

func (s *SessionManagerService) Create(
	ctx context.Context,
	data SessionData,
) (Session, error) {

	session := *NewSession(s.sequence.Generate(), data)

	result, err := s.repository.NonAutoSave(ctx, session)
	if err != nil {
		return NewEmptySession(), err
	}

	return *result, err

}

/* __________________________________________________ */

func (s *SessionManagerService) getSession(
	ctx context.Context,
	user common.UserID,
) (*Session, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(UserIDCondition(user))
	return s.repository.FindOneBySpec(ctx, spec)
}
