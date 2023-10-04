package common

import "context"

/* __________________________________________________ */

type Transactor interface {
	Start(context.Context) (context.Context, error)
	StartEx(context.Context) context.Context
	Close(context.Context, error) error
	CloseEx(context.Context, error)
	Within(context.Context, func(context.Context) error) error
}

/* __________________________________________________ */
