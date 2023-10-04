package user

import (
	"github.com/evgenivanovi/gomart/internal/domain/auth"
	"github.com/evgenivanovi/gomart/internal/domain/auth/session"
	"github.com/evgenivanovi/gomart/internal/domain/auth/token"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
)

/* __________________________________________________ */

type UserData struct {
	Credentials auth.Credentials
}

func NewUserData(credentials auth.Credentials) *UserData {
	return &UserData{
		Credentials: credentials,
	}
}

/* __________________________________________________ */

type User struct {
	id       common.UserID
	data     UserData
	metadata core.Metadata
}

func NewUser(id common.UserID, data UserData, metadata core.Metadata) *User {
	return &User{
		id:       id,
		data:     data,
		metadata: metadata,
	}
}

func NewEmptyUser() User {
	return User{}
}

func NewEmptyUsers() []User {
	return nil
}

func NewEmptyPointerUser() *User {
	return nil
}

func NewEmptyPointerUsers() []*User {
	return nil
}

func ToUserPointers(entities []User) []*User {
	result := make([]*User, 0)
	for _, entity := range entities {
		result = append(result, &entity)
	}
	return result
}

func ToUserValues(entities []*User) []User {
	result := make([]User, 0)
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

func (e *User) Identity() common.UserID {
	return e.id
}

func (e *User) Data() UserData {
	return e.data
}

func (e *User) Metadata() core.Metadata {
	return e.metadata
}

func (e *User) ToAuthUser(data AuthUserData) *AuthUser {
	return NewAuthUser(e.id, data, e.metadata)
}

/* __________________________________________________ */

type AuthUserData struct {
	session.SessionID
	token.Tokens
	UserData
}

func NewAuthUserData(
	sessionID session.SessionID, tokens token.Tokens, userData UserData,
) *AuthUserData {
	return &AuthUserData{
		SessionID: sessionID,
		Tokens:    tokens,
		UserData:  userData,
	}
}

type AuthUser struct {
	id       common.UserID
	data     AuthUserData
	metadata core.Metadata
}

func NewAuthUser(
	id common.UserID, data AuthUserData, metadata core.Metadata,
) *AuthUser {
	return &AuthUser{
		id:       id,
		data:     data,
		metadata: metadata,
	}
}

func NewAuthenticatedEmptyUser() AuthUser {
	return AuthUser{}
}

func NewEmptyAuthUsers() []AuthUser {
	return nil
}

func NewEmptyPointerAuthUser() *AuthUser {
	return nil
}

func NewEmptyPointerAuthUsers() []*AuthUser {
	return nil
}

func (e *AuthUser) Identity() common.UserID {
	return e.id
}

func (e *AuthUser) Data() AuthUserData {
	return e.data
}

func (e *AuthUser) Metadata() core.Metadata {
	return e.metadata
}

func (e *AuthUser) ToUser() *User {
	return NewUser(e.id, e.data.UserData, e.metadata)
}

/* __________________________________________________ */
