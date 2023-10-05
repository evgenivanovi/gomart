package user

import (
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/jackc/pgx/v5"
)

/* __________________________________________________ */

func scanOne(row pgx.Row, record *model.Users) error {
	return row.Scan(
		// ID
		&record.ID,
		// DATA
		&record.Username,
		&record.Password,
		&record.Hashed,
		// METADATA
		&record.CreatedAt,
		&record.UpdatedAt,
		&record.DeletedAt,
	)
}

func scanOneFunc(record *model.Users) func(row pgx.Row) error {
	return func(row pgx.Row) error {
		return scanOne(row, record)
	}
}

func scanMany(rows pgx.Rows, records *[]*model.Users) error {
	defer rows.Close()

	for rows.Next() {
		var record model.Users
		if err := scanOne(rows, &record); err != nil {
			return err
		}
		*records = append(*records, &record)
	}

	return nil
}

func scanManyFunc(records *[]*model.Users) func(rows pgx.Rows) error {
	return func(rows pgx.Rows) error {
		return scanMany(rows, records)
	}
}

/* __________________________________________________ */
