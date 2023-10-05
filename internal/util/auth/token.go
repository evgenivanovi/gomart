package auth

import "time"

/* __________________________________________________ */

const TokenAccessSecretKey = "default"
const TokenAccessExpirationTime = time.Minute * 15

const TokenRefreshSecretKey = "default"
const TokenRefreshExpirationTime = time.Hour * 24 * 7

/* __________________________________________________ */

type TokenOp func(*TokenSettings)

func (o TokenOp) Join(op TokenOp) TokenOp {
	return func(opts *TokenSettings) {
		o(opts)
		op(opts)
	}
}

func (o TokenOp) And(ops ...TokenOp) TokenOp {
	return func(opts *TokenSettings) {
		o(opts)
		for _, fn := range ops {
			fn(opts)
		}
	}
}

func WithAccessSecret(secret string) TokenOp {
	return func(opts *TokenSettings) {
		opts.accessSecret = secret
	}
}

func WithAccessSecretFn(fn func() string) TokenOp {
	return WithAccessSecret(fn())
}

func WithAccessExpiration(expiration time.Duration) TokenOp {
	return func(opts *TokenSettings) {
		opts.accessExpiration = expiration
	}
}

func WithAccessExpirationFn(fn func() time.Duration) TokenOp {
	return WithAccessExpiration(fn())
}

func WithRefreshSecret(secret string) TokenOp {
	return func(opts *TokenSettings) {
		opts.refreshSecret = secret
	}
}

func WithRefreshSecretFn(fn func() string) TokenOp {
	return WithRefreshSecret(fn())
}

func WithRefreshExpiration(expiration time.Duration) TokenOp {
	return func(opts *TokenSettings) {
		opts.refreshExpiration = expiration
	}
}

func WithRefreshExpirationFn(fn func() time.Duration) TokenOp {
	return WithRefreshExpiration(fn())
}

/* __________________________________________________ */

type TokenSettings struct {
	accessSecret      string
	refreshSecret     string
	accessExpiration  time.Duration
	refreshExpiration time.Duration
}

func (t *TokenSettings) AccessSecret() string {
	return t.accessSecret
}

func (t *TokenSettings) RefreshSecret() string {
	return t.refreshSecret
}

func (t *TokenSettings) AccessExpiration() time.Duration {
	return t.accessExpiration
}

func (t *TokenSettings) RefreshExpiration() time.Duration {
	return t.refreshExpiration
}

func NewTokenSettings(ops ...TokenOp) *TokenSettings {
	op := tokenSettings()
	for _, fn := range ops {
		fn(op)
	}
	return op
}

func tokenSettings() *TokenSettings {
	return &TokenSettings{
		accessSecret:      TokenAccessSecretKey,
		refreshSecret:     TokenRefreshSecretKey,
		accessExpiration:  TokenAccessExpirationTime,
		refreshExpiration: TokenRefreshExpirationTime,
	}
}

/* __________________________________________________ */
