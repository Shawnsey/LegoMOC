package daos

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand/v2"

	. "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
	"github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/model"
	"github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/table"
)

type OrderDao interface {
	Insert(ctx context.Context, model model.Order) error
	Delete(ctx context.Context, id int32) error
	Update(ctx context.Context, id int32) (model.Order, error)
}

type OrderPsqlDao struct {
	Client *sql.DB
}

func NewOrderPsqlDao(db *sql.DB) OrderPsqlDao {
	return OrderPsqlDao{Client: db}
}

func (p *OrderPsqlDao) Insert(ctx context.Context, model model.Order) error {
	model.OrderID = rand.Int32()

	insertStmt := table.Order.INSERT(table.Order.OrderID, table.Order.CustomerID, table.Order.LineItems, table.Order.CreatedAt).MODEL(model)

	_, err := insertStmt.ExecContext(ctx, p.Client)
	if err != nil {
		fmt.Errorf("Failed to insert data: %w", err)
		return err
	}
	return nil
}

func (p *OrderPsqlDao) Delete(ctx context.Context, id int32) error {
	deleteStmt := table.Order.DELETE().WHERE(table.Order.OrderID.EQ(Int(int64(id))))

	_, err := deleteStmt.ExecContext(ctx, p.Client)
	if err != nil {
		fmt.Errorf("Entry failed to be deleted %w", err)
		return err
	}
	return nil
}

func (p *OrderPsqlDao) Update(ctx context.Context, id int32) (model.Order, error) {

	return model.Order{}, nil 
}
