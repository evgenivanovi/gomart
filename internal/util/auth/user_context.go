package auth

import (
	"context"
	"net/http"
)

/* __________________________________________________ */

type contextKey string

const ctxKey contextKey = "ctx.user"

/* __________________________________________________ */

// FromCtx
// Takes a context.Context and returns the JWT token associated with it (if any).
func FromCtx(ctx context.Context) *User {
	if value, ok := ctx.Value(ctxKey).(*User); ok {
		return value
	}
	return nil
}

// WithCtx
// Associates a JWT token with a context.Context and returns it.
func WithCtx(ctx context.Context, user *User) context.Context {
	value := FromCtx(ctx)
	if value == user {
		return ctx
	}
	return context.WithValue(ctx, ctxKey, user)
}

func WithRequestCtx(request *http.Request, user *User) {
	ctx := WithCtx(request.Context(), user)
	*request = *request.WithContext(ctx)
}

/* __________________________________________________ */
