package session

import (
	"github.com/evgenivanovi/gomart/internal/domain/auth/token"
	"github.com/evgenivanovi/gomart/internal/domain/common"
)

/* __________________________________________________ */

type SessionID struct {
	id string
}

func NewSessionID(id string) *SessionID {
	return &SessionID{
		id: id,
	}
}

func (s SessionID) ID() string {
	return s.id
}

/* __________________________________________________ */

type SessionData struct {
	UserID common.UserID
	Token  token.RefreshToken
}

func NewSessionData(userID common.UserID, token token.RefreshToken) *SessionData {
	return &SessionData{
		UserID: userID,
		Token:  token,
	}
}

/* __________________________________________________ */

type Session struct {
	id   SessionID
	data SessionData
}

func NewSession(id SessionID, data SessionData) *Session {
	return &Session{
		id:   id,
		data: data,
	}
}

func NewEmptySession() Session {
	return Session{}
}

func NewEmptyPointerSession() *Session {
	return nil
}

func NewEmptySessions() []Session {
	return nil
}

func NewEmptyPointerSessions() []*Session {
	return nil
}

func ToSessionPointers(entities []Session) []*Session {
	result := make([]*Session, 0)
	for _, entity := range entities {
		result = append(result, &entity)
	}
	return result
}

func ToSessionValues(entities []*Session) []Session {
	result := make([]Session, 0)
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

func (e *Session) Identity() SessionID {
	return e.id
}

func (e *Session) Data() SessionData {
	return e.data
}

/* __________________________________________________ */
