//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Withdrawals = newWithdrawalsTable("public", "withdrawals", "")

type withdrawalsTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	UserID    postgres.ColumnInteger
	Order     postgres.ColumnString
	Amount    postgres.ColumnFloat
	CreatedAt postgres.ColumnTimestampz
	UpdatedAt postgres.ColumnTimestampz
	DeletedAt postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type WithdrawalsTable struct {
	withdrawalsTable

	EXCLUDED withdrawalsTable
}

// AS creates new WithdrawalsTable with assigned alias
func (a WithdrawalsTable) AS(alias string) *WithdrawalsTable {
	return newWithdrawalsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new WithdrawalsTable with assigned schema name
func (a WithdrawalsTable) FromSchema(schemaName string) *WithdrawalsTable {
	return newWithdrawalsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new WithdrawalsTable with assigned table prefix
func (a WithdrawalsTable) WithPrefix(prefix string) *WithdrawalsTable {
	return newWithdrawalsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new WithdrawalsTable with assigned table suffix
func (a WithdrawalsTable) WithSuffix(suffix string) *WithdrawalsTable {
	return newWithdrawalsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newWithdrawalsTable(schemaName, tableName, alias string) *WithdrawalsTable {
	return &WithdrawalsTable{
		withdrawalsTable: newWithdrawalsTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newWithdrawalsTableImpl("", "excluded", ""),
	}
}

func newWithdrawalsTableImpl(schemaName, tableName, alias string) withdrawalsTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		UserIDColumn    = postgres.IntegerColumn("user_id")
		OrderColumn     = postgres.StringColumn("order")
		AmountColumn    = postgres.FloatColumn("amount")
		CreatedAtColumn = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn = postgres.TimestampzColumn("updated_at")
		DeletedAtColumn = postgres.TimestampzColumn("deleted_at")
		allColumns      = postgres.ColumnList{IDColumn, UserIDColumn, OrderColumn, AmountColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
		mutableColumns  = postgres.ColumnList{UserIDColumn, OrderColumn, AmountColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
	)

	return withdrawalsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		UserID:    UserIDColumn,
		Order:     OrderColumn,
		Amount:    AmountColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		DeletedAt: DeletedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}