package session

import (
	"fmt"

	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/table"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
)

/* __________________________________________________ */

func insertOneStatement(
	record model.Sessions,
) (string, []interface{}) {

	stmt := table.Sessions.
		INSERT(
			// ID
			table.Sessions.ID,
			// DATA
			table.Sessions.UserID,
			table.Sessions.Token,
			table.Sessions.ExpiresAt,
		).
		VALUES(
			// ID
			record.ID,
			// DATA
			record.UserID,
			record.Token,
			record.ExpiresAt,
		)

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

func insertAllStatement(
	records []model.Sessions,
) (string, []interface{}) {

	stmt := table.Sessions.INSERT(
		// ID
		table.Sessions.ID,
		// DATA
		table.Sessions.UserID,
		table.Sessions.Token,
		table.Sessions.ExpiresAt,
	)

	for _, record := range records {
		stmt = stmt.VALUES(
			// ID
			record.ID,
			// DATA
			record.UserID,
			record.Token,
			record.ExpiresAt,
		)
	}

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

/* __________________________________________________ */
