package session

import (
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/jackc/pgx/v5"
)

/* __________________________________________________ */

func scanOne(row pgx.Row, record *model.Sessions) error {
	return row.Scan(
		// ID
		&record.ID,
		// DATA
		&record.UserID,
		&record.Token,
		&record.ExpiresAt,
	)
}

func scanOneFunc(record *model.Sessions) func(row pgx.Row) error {
	return func(row pgx.Row) error {
		return scanOne(row, record)
	}
}

func scanMany(rows pgx.Rows, records *[]*model.Sessions) error {
	defer rows.Close()

	for rows.Next() {
		var record model.Sessions
		if err := scanOne(rows, &record); err != nil {
			return err
		}
		*records = append(*records, &record)
	}

	return nil
}

func scanManyFunc(records *[]*model.Sessions) func(rows pgx.Rows) error {
	return func(rows pgx.Rows) error {
		return scanMany(rows, records)
	}
}

/* __________________________________________________ */
