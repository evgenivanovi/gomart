package session

import (
	"fmt"

	"github.com/evgenivanovi/gomart/internal/domain/auth/session"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/table"
	"github.com/evgenivanovi/gomart/pkg/search"
	slices "github.com/evgenivanovi/gomart/pkg/std/slice"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
	pgjet "github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

func idExpression(
	id session.SessionID,
) pgjet.BoolExpression {
	return table.Sessions.ID.EQ(pgjet.String(id.ID()))
}

func idsExpression(
	ids []session.SessionID,
) pgjet.BoolExpression {
	exp := make([]pgjet.Expression, 0)
	for _, id := range ids {
		exp = append(exp, pgjet.String(id.ID()))
	}
	return table.Sessions.ID.IN(exp...)
}

func userExpression(
	id common.UserID,
) pgjet.BoolExpression {
	return table.Sessions.UserID.EQ(pgjet.Int(id.ID()))
}

func usersExpression(
	ids []common.UserID,
) pgjet.BoolExpression {
	exp := make([]pgjet.Expression, 0)
	for _, id := range ids {
		exp = append(exp, pgjet.Int(id.ID()))
	}
	return table.Sessions.UserID.IN(exp...)
}

/* __________________________________________________ */

func buildQuery(
	searchExp pgjet.BoolExpression,
	orderExp []pgjet.OrderByClause,
	lock pgjet.RowLock,
	sliceCondition search.SliceCondition,
) (string, []interface{}) {

	stmt := pgjet.
		SELECT(
			table.Sessions.AllColumns,
		).
		FROM(
			table.Sessions,
		)

	if sliceCondition.Chunked() {
		chunk := sliceCondition.Chunk()
		if limit, ok := chunk.Limit(); ok {
			stmt.LIMIT(limit)
		}
		if offset, ok := chunk.Offset(); ok {
			stmt.OFFSET(offset)
		}
	}

	if searchExp != nil {
		stmt = stmt.WHERE(searchExp)
	}

	if slices.IsNotEmpty(orderExp) {
		stmt = stmt.ORDER_BY(orderExp...)
	}

	if lock != nil {
		stmt = stmt.FOR(lock)
	}

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

/* __________________________________________________ */
