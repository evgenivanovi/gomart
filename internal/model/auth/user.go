package auth

import (
	"github.com/evgenivanovi/gomart/internal/model"
)

/* __________________________________________________ */

type UserData struct {
	Credentials Credentials
}

type User struct {
	ID       int64
	Data     UserData
	Metadata model.Metadata
}

/* __________________________________________________ */
