package user

import (
	"fmt"

	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/table"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
)

/* __________________________________________________ */

func insertOneStatement(
	record model.Users,
) (string, []interface{}) {

	stmt := table.Users.
		INSERT(
			// DATA
			table.Users.Username,
			table.Users.Password,
			table.Users.Hashed,
			// METADATA
			table.Users.CreatedAt,
			table.Users.UpdatedAt,
			table.Users.DeletedAt,
		).
		VALUES(
			// DATA
			record.Username,
			record.Password,
			record.Hashed,
			// METADATA
			record.CreatedAt,
			record.UpdatedAt,
			record.DeletedAt,
		).
		RETURNING(
			table.Users.ID,
		)

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

func insertAllStatement(
	records []model.Users,
) (string, []interface{}) {

	stmt := table.Users.INSERT(
		// ID
		table.Users.ID,
		// DATA
		table.Users.Username,
		table.Users.Password,
		table.Users.Hashed,
		// METADATA
		table.Users.CreatedAt,
		table.Users.UpdatedAt,
		table.Users.DeletedAt,
	)

	for _, record := range records {
		stmt = stmt.VALUES(
			// ID
			record.ID,
			// DATA
			record.Username,
			record.Password,
			record.Hashed,
			// METADATA
			record.CreatedAt,
			record.UpdatedAt,
			record.DeletedAt,
		)
	}

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

/* __________________________________________________ */
