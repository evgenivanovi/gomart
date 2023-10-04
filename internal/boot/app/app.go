package app

import (
	"flag"

	"github.com/evgenivanovi/gomart/pkg/fw"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
)

/* __________________________________________________ */

var (
	/* Framework Properties */
	Application fw.Application

	/* Server Properties */
	ServerAddressProperty fw.Property

	/* Authentication Properties */
	TokenAccessSecretKeyProperty      fw.Property
	TokenAccessTimeExpirationProperty fw.Property

	TokenRefreshSecretKeyProperty      fw.Property
	TokenRefreshTimeExpirationProperty fw.Property

	/* Postgres Properties */
	DSNPostgresProperty fw.Property

	/* Accrual Properties */
	AccrualProperty fw.Property
)

/* __________________________________________________ */

/* Server Properties */
const HTTPServerAddressProp = "server.address"
const HTTPServerAddressEnvVar = "RUN_ADDRESS"
const HTTPServerAddressArgVar = "a"
const HTTPServerAddressDefaultValue = "localhost:8080"

/* Authentication Properties */
const TokenAccessSecretKeyProp = "token.access.secret"
const TokenAccessSecretKeyEnvVar = "TOKEN_ACCESS_SECRET_KEY"
const TokenAccessSecretKeyArgVar = "token-access-secret-key"

const TokenRefreshSecretKeyProp = "token.refresh.secret"
const TokenRefreshSecretKeyEnvVar = "TOKEN_REFRESH_SECRET_KEY"
const TokenRefreshSecretKeyArgVar = "token-refresh-secret-key"

const TokenTimeExpirationKeyProp = "token.time.expiration"
const TokenTimeExpirationKeyEnvVar = "TOKEN_TIME_EXPIRATION"
const TokenTimeExpirationKeyArgVar = "token-time-expiration"

/* Postgres Properties */
const DSNPostgresProp = "dsn"
const DSNPostgresEnvVar = "DATABASE_URI"
const DSNPostgresArgVar = "d"

/* Accrual Properties */
const AccrualProp = "accrual.address"
const AccrualEnvVar = "ACCRUAL_SYSTEM_ADDRESS"
const AccrualArgVar = "r"

/* __________________________________________________ */

func init() {

	ServerAddressProperty = fw.NewProperty(HTTPServerAddressProp)
	ServerAddressProperty.BindOne(fw.NewEnvSource(HTTPServerAddressEnvVar))
	ServerAddressProperty.BindOne(fw.NewArgSource(HTTPServerAddressArgVar))

	TokenAccessSecretKeyProperty = fw.NewProperty(TokenAccessSecretKeyProp)
	TokenAccessSecretKeyProperty.BindOne(fw.NewEnvSource(TokenAccessSecretKeyEnvVar))
	TokenAccessSecretKeyProperty.BindOne(fw.NewArgSource(TokenAccessSecretKeyArgVar))

	TokenRefreshSecretKeyProperty = fw.NewProperty(TokenRefreshSecretKeyProp)
	TokenRefreshSecretKeyProperty.BindOne(fw.NewEnvSource(TokenRefreshSecretKeyEnvVar))
	TokenRefreshSecretKeyProperty.BindOne(fw.NewArgSource(TokenRefreshSecretKeyArgVar))

	TokenRefreshTimeExpirationProperty = fw.NewProperty(TokenTimeExpirationKeyProp)
	TokenRefreshTimeExpirationProperty.BindOne(fw.NewEnvSource(TokenTimeExpirationKeyEnvVar))
	TokenRefreshTimeExpirationProperty.BindOne(fw.NewArgSource(TokenTimeExpirationKeyArgVar))

	DSNPostgresProperty = fw.NewProperty(DSNPostgresProp)
	DSNPostgresProperty.BindOne(fw.NewEnvSource(DSNPostgresEnvVar))
	DSNPostgresProperty.BindOne(fw.NewArgSource(DSNPostgresArgVar))

	AccrualProperty = fw.NewProperty(AccrualProp)
	AccrualProperty.BindOne(fw.NewEnvSource(AccrualEnvVar))
	AccrualProperty.BindOne(fw.NewArgSource(AccrualArgVar))

	flag.Parse()

}

/* __________________________________________________ */

func BootApp() {

	Application = *fw.NewApplication()

	serverAddressFn := fw.FirstStringNotEmpty(HTTPServerAddressDefaultValue)
	serverAddress := ServerAddressProperty.Calc(serverAddressFn)
	Application.ServerOpts = *fw.NewServerOpts(
		fw.WithAddress(serverAddress),
	)

	slogx.Log().Debug(
		"Calculated server options: " + Application.ServerOpts.Address(),
	)

}

/* __________________________________________________ */
