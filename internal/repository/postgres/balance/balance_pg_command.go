package balance

import (
	"fmt"

	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/table"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
	"github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

func insertOneStatement(
	record model.Balances,
) (string, []interface{}) {

	stmt := table.Balances.
		INSERT(
			table.Balances.UserID,
			table.Balances.Amount,
		).
		VALUES(
			record.UserID,
			record.Amount,
		).
		RETURNING(
			table.Balances.ID,
		)

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

func updateOneStatement(
	record model.Balances,
) (string, []interface{}) {

	stmt := table.Balances.
		UPDATE(
			table.Balances.UserID,
			table.Balances.Amount,
		)

	stmt = stmt.SET(
		record.UserID,
		record.Amount,
	)

	stmt = stmt.WHERE(
		table.Balances.ID.EQ(postgres.Int(record.ID)),
	)

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

/* __________________________________________________ */
