package token

import (
	"time"

	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/util/auth"
	"github.com/golang-jwt/jwt/v5"
)

/* __________________________________________________ */

func CreateAccessToken(data TokenData, expiration time.Duration) *jwt.Token {
	return auth.CreateAccessToken(MapUserFrom(data), expiration)
}

func CreateRefreshToken(data TokenData, expiration time.Duration) *jwt.Token {
	return auth.CreateRefreshToken(MapUserFrom(data), expiration)
}

/* __________________________________________________ */

// MapUserFrom
// The function has the prefix From, because in pure or hexagonal architecture everything starts with the domain.
// Therefore, From suffix has an initial relation to the domain entity.
func MapUserFrom(data TokenData) auth.User {
	return *auth.NewUser(data.UserID.ID())
}

// MapUserTo
// The function has the prefix To, because in pure or hexagonal architecture everything starts with the domain.
// Therefore, To suffix has an initial relation to the domain entity.
func MapUserTo(entity auth.User) TokenData {
	return TokenData{
		UserID: *common.NewUserID(entity.UserID),
	}
}

/* __________________________________________________ */
