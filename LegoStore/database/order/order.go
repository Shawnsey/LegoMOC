package order

import (
	"context"
	"database/sql"
	"fmt"

	. "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
	"github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/model"
	"github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/table"
)

type Postgresdb struct {
	Client sql.DB
}

func (p *Postgresdb) Insert(ctx context.Context, order model.Order) error {

	insertStmt := table.Order.INSERT(table.Order.OrderID, table.Order.CustomerID, table.Order.LineItems, table.Order.CreatedAt).MODEL(order)

	_, err := insertStmt.ExecContext(ctx, &p.Client)
	if err != nil {
		fmt.Errorf("Failed to insert data: %w", err)
		return err
	}
	return nil
}

func (p *Postgresdb) Delete(ctx context.Context, id int32) error {
	// // deleteStmt := table.Order.DELETE().WHERE(table.Order.OrderID.EQ(Int(int64(id))))
	// deleteStmt := table.Order.DELETE().WHERE(table.Order.OrderID.EQ(Int(int64(id))))

	// _, err := deleteStmt.ExecContext(ctx, &p.Client)
	// if err != nil {
	// 	fmt.Errorf("Entry failed to be deleted %w", err)
	// 	return err
	// }
	return nil
}
