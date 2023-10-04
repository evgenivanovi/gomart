package withdraw

import (
	"fmt"

	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/table"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
	"github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

func insertOneStatement(
	record model.Withdraws,
) (string, []interface{}) {

	stmt := table.Withdraws.
		INSERT(
			table.Withdraws.UserID,
			table.Withdraws.Amount,
		).
		VALUES(
			record.UserID,
			record.Amount,
		).
		RETURNING(
			table.Withdraws.ID,
		)

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

func updateOneStatement(
	record model.Withdraws,
) (string, []interface{}) {

	stmt := table.Withdraws.
		UPDATE(
			table.Withdraws.UserID,
			table.Withdraws.Amount,
		)

	stmt = stmt.SET(
		record.UserID,
		record.Amount,
	)

	stmt = stmt.WHERE(
		table.Withdraws.ID.EQ(postgres.Int(record.ID)),
	)

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

/* __________________________________________________ */
