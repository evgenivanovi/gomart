package withdraw

import (
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/jackc/pgx/v5"
)

/* __________________________________________________ */

func scanOne(row pgx.Row, record *model.Withdraws) error {
	return row.Scan(
		&record.ID,
		&record.UserID,
		&record.Amount,
	)
}

func scanOneFunc(record *model.Withdraws) func(row pgx.Row) error {
	return func(row pgx.Row) error {
		return scanOne(row, record)
	}
}

func scanMany(rows pgx.Rows, records *[]*model.Withdraws) error {
	defer rows.Close()

	for rows.Next() {
		var record model.Withdraws
		if err := scanOne(rows, &record); err != nil {
			return err
		}
		*records = append(*records, &record)
	}

	return nil
}

func scanManyFunc(records *[]*model.Withdraws) func(rows pgx.Rows) error {
	return func(rows pgx.Rows) error {
		return scanMany(rows, records)
	}
}

/* __________________________________________________ */
