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

var Balances = newBalancesTable("public", "balances", "")

type balancesTable struct {
	postgres.Table

	// Columns
	ID     postgres.ColumnInteger
	UserID postgres.ColumnInteger
	Amount postgres.ColumnFloat

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type BalancesTable struct {
	balancesTable

	EXCLUDED balancesTable
}

// AS creates new BalancesTable with assigned alias
func (a BalancesTable) AS(alias string) *BalancesTable {
	return newBalancesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new BalancesTable with assigned schema name
func (a BalancesTable) FromSchema(schemaName string) *BalancesTable {
	return newBalancesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new BalancesTable with assigned table prefix
func (a BalancesTable) WithPrefix(prefix string) *BalancesTable {
	return newBalancesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new BalancesTable with assigned table suffix
func (a BalancesTable) WithSuffix(suffix string) *BalancesTable {
	return newBalancesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newBalancesTable(schemaName, tableName, alias string) *BalancesTable {
	return &BalancesTable{
		balancesTable: newBalancesTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newBalancesTableImpl("", "excluded", ""),
	}
}

func newBalancesTableImpl(schemaName, tableName, alias string) balancesTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		UserIDColumn   = postgres.IntegerColumn("user_id")
		AmountColumn   = postgres.FloatColumn("amount")
		allColumns     = postgres.ColumnList{IDColumn, UserIDColumn, AmountColumn}
		mutableColumns = postgres.ColumnList{UserIDColumn, AmountColumn}
	)

	return balancesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:     IDColumn,
		UserID: UserIDColumn,
		Amount: AmountColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
