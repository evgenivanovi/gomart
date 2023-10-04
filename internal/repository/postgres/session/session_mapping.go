package session

import (
	"github.com/evgenivanovi/gomart/internal/domain/auth/session"
	"github.com/evgenivanovi/gomart/internal/domain/auth/token"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
)

/* __________________________________________________ */

func ToSession(record *model.Sessions) *session.Session {

	id := session.NewSessionID(
		record.ID,
	)

	data := session.NewSessionData(
		*common.NewUserID(record.UserID),
		*token.NewRefreshToken(record.Token, record.ExpiresAt),
	)

	return session.NewSession(*id, *data)

}

func FromSession(entity session.Session) model.Sessions {
	return model.Sessions{
		ID:        entity.Identity().ID(),
		UserID:    entity.Data().UserID.ID(),
		Token:     entity.Data().Token.Token,
		ExpiresAt: entity.Data().Token.ExpiresAt,
	}
}

func ToSessions(records []*model.Sessions) []*session.Session {
	entities := make([]*session.Session, 0)
	for _, record := range records {
		entities = append(entities, ToSession(record))
	}
	return entities
}

func FromSessions(entities []session.Session) []model.Sessions {
	records := make([]model.Sessions, 0)
	for _, entity := range entities {
		records = append(records, FromSession(entity))
	}
	return records
}

/* __________________________________________________ */
