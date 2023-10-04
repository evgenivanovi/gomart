package auth

import (
	pg_infra "github.com/evgenivanovi/gomart/internal/boot/pg"
	session_dm "github.com/evgenivanovi/gomart/internal/domain/auth/session"
	session_pg "github.com/evgenivanovi/gomart/internal/repository/postgres/session"
)

/* __________________________________________________ */

var (
	SessionReadRepository  session_dm.SessionReadRepository
	SessionWriteRepository session_dm.SessionWriteRepository
	SessionRepository      session_dm.SessionRepository

	SessionIDGenerator session_dm.SessionIDGenerator
	SessionManager     session_dm.SessionManager
)

/* __________________________________________________ */

func BootAuthSession() {

	SessionReadRepository = session_pg.ProvideSessionReadRepositoryService(
		pg_infra.PostgresReadRequester,
	)
	SessionWriteRepository = session_pg.ProvideSessionWriteRepositoryService(
		pg_infra.PostgresWriteRequester,
	)
	SessionRepository = session_dm.ProvideSessionRepositoryService(
		SessionReadRepository, SessionWriteRepository,
	)

	SessionIDGenerator = session_dm.ProvideSessionIDGeneratorService()

	SessionManager = session_dm.ProvideSessionManagerService(
		SessionIDGenerator, SessionRepository,
	)

}

/* __________________________________________________ */
