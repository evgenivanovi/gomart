package withdrawal

import (
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/jackc/pgx/v5"
)

/* __________________________________________________ */

func scanOne(row pgx.Row, record *model.Withdrawals) error {
	return row.Scan(
		// ID
		&record.ID,
		// DATA
		&record.UserID,
		&record.Order,
		&record.Amount,
		// METADATA
		&record.CreatedAt,
		&record.UpdatedAt,
		&record.DeletedAt,
	)
}

func scanOneFunc(record *model.Withdrawals) func(row pgx.Row) error {
	return func(row pgx.Row) error {
		return scanOne(row, record)
	}
}

func scanMany(rows pgx.Rows, records *[]*model.Withdrawals) error {
	defer rows.Close()

	for rows.Next() {
		var record model.Withdrawals
		if err := scanOne(rows, &record); err != nil {
			return err
		}
		*records = append(*records, &record)
	}

	return nil
}

func scanManyFunc(records *[]*model.Withdrawals) func(rows pgx.Rows) error {
	return func(rows pgx.Rows) error {
		return scanMany(rows, records)
	}
}

/* __________________________________________________ */
