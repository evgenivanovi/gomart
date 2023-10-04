package withdrawal

import (
	"fmt"

	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/table"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
)

/* __________________________________________________ */

func insertOneStatement(
	record model.Withdrawals,
) (string, []interface{}) {

	stmt := table.Withdrawals.
		INSERT(
			table.Withdrawals.UserID,
			table.Withdrawals.Order,
			table.Withdrawals.Amount,
			table.Withdrawals.CreatedAt,
			table.Withdrawals.UpdatedAt,
			table.Withdrawals.DeletedAt,
		).
		VALUES(
			// DATA
			record.UserID,
			record.Order,
			record.Amount,
			// METADATA
			record.CreatedAt,
			record.UpdatedAt,
			record.DeletedAt,
		).
		RETURNING(
			table.Withdrawals.ID,
		)

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

/* __________________________________________________ */
