package auth

import (
	"fmt"
	"net/http"
	"time"

	timex "github.com/evgenivanovi/gomart/pkg/std/time"
	"github.com/evgenivanovi/gomart/pkg/stdx/jwtx"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
	"github.com/evgenivanovi/gomart/pkg/stdx/net/http/headers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"github.com/pkg/errors"
)

/* __________________________________________________ */

const CookieAuthKey = "auth"

/* __________________________________________________ */

type AccessClaims struct {
	jwt.RegisteredClaims
	User *User `json:"user"`
}

type RefreshClaims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}

/* __________________________________________________ */

func KeyProvider(secret string) func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}
}

func MethodProvider() jwt.SigningMethod {
	return jwt.SigningMethodHS256
}

func ClaimsProvider() func() jwt.Claims {
	return func() jwt.Claims {
		return &AccessClaims{}
	}
}

func ExtractorProvider() request.Extractor {
	return request.MultiExtractor{
		&request.PostExtractionFilter{
			Extractor: request.HeaderExtractor{
				headers.AuthorizationKey.String(),
			},
			Filter: jwtx.ExtractBearerToken,
		},
	}
}

func AfterProvider() func(http.ResponseWriter, *http.Request, *jwt.Token, string) error {
	return func(httpWriter http.ResponseWriter, httpRequest *http.Request, token *jwt.Token, tokenString string) error {

		if token == nil || !token.Valid {
			httpWriter.WriteHeader(http.StatusUnauthorized)
			return nil
		}

		var claims *AccessClaims
		if tokenClaims, ok := token.Claims.(*AccessClaims); ok {
			claims = tokenClaims
		}

		if claims == nil || claims.User == nil || claims.User.UserID == 0 {
			httpWriter.WriteHeader(http.StatusUnauthorized)
			return nil
		}

		if time.Now().After(claims.ExpiresAt.Time) {
			httpWriter.WriteHeader(http.StatusUnauthorized)
			return nil
		}

		WithRequestCtx(httpRequest, claims.User)
		WriteTokenToResponseHeader(httpWriter, token, tokenString)

		return nil
	}
}

func RecoveryProvider() func(http.ResponseWriter, *http.Request, error) {
	return func(httpWriter http.ResponseWriter, httpRequest *http.Request, err error) {
		if errors.Is(err, request.ErrNoTokenInRequest) {
			slogx.Log().Debug(fmt.Sprintf("authentication: %v", err))
		}
		httpWriter.WriteHeader(http.StatusUnauthorized)
	}
}

/* __________________________________________________ */

func WriteTokenToRequestContext(request *http.Request, token *jwt.Token, tokenString string) {
	ctx := request.Context()
	if token != nil {
		ctx = jwtx.WithCtx(ctx, token)
	}
	if tokenString != "" {
		ctx = jwtx.WithCtxAsString(ctx, tokenString)
	}
	*request = *request.WithContext(ctx)
}

func WriteTokenToRequestCookie(request *http.Request, token *jwt.Token, tokenString string) {

	tokenExpirationTime, _ := token.Claims.GetExpirationTime()
	cookieExpirationTime := int(tokenExpirationTime.Unix())

	cookie := &http.Cookie{
		Name:   CookieAuthKey,
		Value:  tokenString,
		MaxAge: cookieExpirationTime,
	}

	request.AddCookie(cookie)

}

func WriteTokenToResponseCookie(writer http.ResponseWriter, token *jwt.Token, tokenString string) {

	tokenExpirationTime, _ := token.Claims.GetExpirationTime()
	cookieExpirationTime := int(tokenExpirationTime.Unix())

	cookie := &http.Cookie{
		Name:   CookieAuthKey,
		Value:  tokenString,
		MaxAge: cookieExpirationTime,
	}

	http.SetCookie(writer, cookie)

}

func WriteTokenToRequestHeader(request *http.Request, token *jwt.Token, tokenString string) {
	request.Header.Add(
		headers.AuthorizationKey.String(),
		headers.BuildBearerToken(tokenString),
	)
}

func WriteTokenToResponseHeader(writer http.ResponseWriter, token *jwt.Token, tokenString string) {
	writer.Header().Add(
		headers.AuthorizationKey.String(),
		headers.BuildBearerToken(tokenString),
	)
}

/* __________________________________________________ */

func CreateAccessToken(user User, expiration time.Duration) *jwt.Token {

	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		AccessClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(
					timex.NowPlus(expiration),
				),
			},
			User: &user,
		},
	)

}

func CreateRefreshToken(user User, expiration time.Duration) *jwt.Token {

	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		RefreshClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(
					timex.NowPlus(expiration),
				),
			},
			UserID: user.UserID,
		},
	)

}

/* __________________________________________________ */

func DecodeClaims(token string, secret string) (*AccessClaims, error) {

	claims := &AccessClaims{}

	tkn, err := jwt.ParseWithClaims(
		token,
		claims,
		KeyProvider(secret),
	)

	if err != nil && !tkn.Valid {
		return nil, err
	}

	return claims, nil

}

func DecodeUser(token string, secret string) (*User, error) {

	claims, err := DecodeClaims(token, secret)

	if err != nil {
		return nil, err
	}

	return claims.User, nil

}

/* __________________________________________________ */
