package pg

import (
	"github.com/evgenivanovi/gomart/pkg/goose"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
)

/* __________________________________________________ */

func BootPGMigrations() {

	if !IsInitializedPG() {
		return
	}

	slogx.Log().Debug(
		"Starting PostgreSQL migrations.",
	)

	goose.MigrateUp("./migrations", "postgres", DSN)

	slogx.Log().Debug(
		"Finished PostgreSQL migrations.",
	)

}

/* __________________________________________________ */
