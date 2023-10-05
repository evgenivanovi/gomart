package balance

import (
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/jackc/pgx/v5"
)

/* __________________________________________________ */

func scanOne(row pgx.Row, record *model.Balances) error {
	return row.Scan(
		&record.ID,
		&record.UserID,
		&record.Amount,
	)
}

func scanOneFunc(record *model.Balances) func(row pgx.Row) error {
	return func(row pgx.Row) error {
		return scanOne(row, record)
	}
}

func scanMany(rows pgx.Rows, records *[]*model.Balances) error {
	defer rows.Close()

	for rows.Next() {
		var record model.Balances
		if err := scanOne(rows, &record); err != nil {
			return err
		}
		*records = append(*records, &record)
	}

	return nil
}

func scanManyFunc(records *[]*model.Balances) func(rows pgx.Rows) error {
	return func(rows pgx.Rows) error {
		return scanMany(rows, records)
	}
}

/* __________________________________________________ */
