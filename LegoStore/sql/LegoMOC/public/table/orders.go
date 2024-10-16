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

var Orders = newOrdersTable("public", "orders", "")

type ordersTable struct {
	postgres.Table

	// Columns
	OrderID     postgres.ColumnString
	CustomerID  postgres.ColumnString
	LineItems   postgres.ColumnString
	CreatedAt   postgres.ColumnTimestampz
	ShippedAt   postgres.ColumnTimestampz
	CompletedAt postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrdersTable struct {
	ordersTable

	EXCLUDED ordersTable
}

// AS creates new OrdersTable with assigned alias
func (a OrdersTable) AS(alias string) *OrdersTable {
	return newOrdersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrdersTable with assigned schema name
func (a OrdersTable) FromSchema(schemaName string) *OrdersTable {
	return newOrdersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrdersTable with assigned table prefix
func (a OrdersTable) WithPrefix(prefix string) *OrdersTable {
	return newOrdersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrdersTable with assigned table suffix
func (a OrdersTable) WithSuffix(suffix string) *OrdersTable {
	return newOrdersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrdersTable(schemaName, tableName, alias string) *OrdersTable {
	return &OrdersTable{
		ordersTable: newOrdersTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newOrdersTableImpl("", "excluded", ""),
	}
}

func newOrdersTableImpl(schemaName, tableName, alias string) ordersTable {
	var (
		OrderIDColumn     = postgres.StringColumn("order_id")
		CustomerIDColumn  = postgres.StringColumn("customer_id")
		LineItemsColumn   = postgres.StringColumn("line_items")
		CreatedAtColumn   = postgres.TimestampzColumn("created_at")
		ShippedAtColumn   = postgres.TimestampzColumn("shipped_at")
		CompletedAtColumn = postgres.TimestampzColumn("completed_at")
		allColumns        = postgres.ColumnList{OrderIDColumn, CustomerIDColumn, LineItemsColumn, CreatedAtColumn, ShippedAtColumn, CompletedAtColumn}
		mutableColumns    = postgres.ColumnList{CustomerIDColumn, LineItemsColumn, CreatedAtColumn, ShippedAtColumn, CompletedAtColumn}
	)

	return ordersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		OrderID:     OrderIDColumn,
		CustomerID:  CustomerIDColumn,
		LineItems:   LineItemsColumn,
		CreatedAt:   CreatedAtColumn,
		ShippedAt:   ShippedAtColumn,
		CompletedAt: CompletedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
