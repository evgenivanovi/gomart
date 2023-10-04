package user

import (
	"fmt"

	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/table"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
	"github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

func insertOneStatement(
	record model.Orders,
) (string, []interface{}) {

	stmt := table.Orders.
		INSERT(
			// ID
			table.Orders.ID,
			// DATA
			table.Orders.UserID,
			table.Orders.Status,
			table.Orders.Accrual,
			// METADATA
			table.Orders.CreatedAt,
			table.Orders.UpdatedAt,
			table.Orders.DeletedAt,
		).
		VALUES(
			// ID
			record.ID,
			// DATA
			record.UserID,
			record.Status,
			record.Accrual,
			// METADATA
			record.CreatedAt,
			record.UpdatedAt,
			record.DeletedAt,
		)

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

func insertAllStatement(
	records []model.Orders,
) (string, []interface{}) {

	stmt := table.Orders.INSERT(
		// ID
		table.Orders.ID,
		// DATA
		table.Orders.UserID,
		table.Orders.Status,
		table.Orders.Accrual,
		// METADATA
		table.Orders.CreatedAt,
		table.Orders.UpdatedAt,
		table.Orders.DeletedAt,
	)

	for _, record := range records {
		stmt = stmt.VALUES(
			// ID
			record.ID,
			// DATA
			record.UserID,
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

func updateOneStatement(
	record model.Orders,
) (string, []interface{}) {

	stmt := table.Orders.
		UPDATE(
			// ID
			table.Orders.ID,
			// DATA
			table.Orders.UserID,
			table.Orders.Status,
			table.Orders.Accrual,
		)

	stmt = stmt.SET(
		// ID
		record.ID,
		// DATA
		record.UserID,
		record.Status,
		record.Accrual,
	)

	stmt = stmt.WHERE(
		table.Orders.ID.EQ(postgres.String(record.ID)),
	)

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

/* __________________________________________________ */
