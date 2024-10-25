package daos

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/model"
	. "github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/table"
)

type OrderDao interface {
	Insert(ctx context.Context, model model.Orders) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, update OrderUpdateBody) (model.Orders, error)
	List(ctx context.Context, id uuid.UUID) ([]struct{model.Orders}, error)
	GetById(ctx context.Context, id uuid.UUID) (struct{model.Orders}, error)
}

type OrderPsqlDao struct {
	Client *sql.DB
}

type OrderUpdateBody struct {
	OrderId uuid.UUID
	ShippedAt  time.Time `json:"shipped_at"`
	CompletedAt time.Time `json:"completed_at"`
}

func NewOrderPsqlDao(db *sql.DB) *OrderPsqlDao {
	return &OrderPsqlDao{Client: db}
}
// create works
func (p *OrderPsqlDao) Insert(ctx context.Context, model model.Orders) error {

	insertStmt := Orders.INSERT(Orders.OrderID, Orders.CustomerID, Orders.LineItems, Orders.CreatedAt).MODEL(model)

	_, err := insertStmt.ExecContext(ctx, p.Client)
	if err != nil {
		fmt.Errorf("Failed to insert data: %w", err)
		return err
	}
	return nil
}
// Delete works
func (p *OrderPsqlDao) Delete(ctx context.Context, id uuid.UUID) error {
	deleteStmt := Orders.DELETE().WHERE(Orders.OrderID.IN(jet.UUID(id)))

	_, err := deleteStmt.ExecContext(ctx, p.Client)
	if err != nil {
		fmt.Errorf("Entry failed to be deleted %w", err)
		return err
	}
	return nil
}

func (p *OrderPsqlDao) GetById(ctx context.Context, id uuid.UUID) (struct{model.Orders}, error) {
	var res struct {
		model.Orders
	}
	getStmt := Orders.SELECT(Orders.AllColumns).FROM(Orders).WHERE(Orders.OrderID.EQ(jet.UUID(id)))
	err := getStmt.QueryContext(ctx, p.Client, &res)
	if err != nil {
		fmt.Errorf("Failed when executing query %w", err)
		return struct{model.Orders}{}, err
	}

	return res, nil 
}
// update works
func (p *OrderPsqlDao) Update(ctx context.Context, update OrderUpdateBody) (model.Orders, error) {
	var updateStmt jet.UpdateStatement
	if (!update.ShippedAt.IsZero() && !update.CompletedAt.IsZero()){
		updateStmt = Orders.UPDATE(Orders.ShippedAt, Orders.CompletedAt).
		SET(update.ShippedAt, update.CompletedAt).
		WHERE(Orders.OrderID.EQ(jet.UUID(update.OrderId))).RETURNING(Orders.AllColumns)
	} else if !update.ShippedAt.IsZero() {
		updateStmt = Orders.UPDATE(Orders.ShippedAt).
		SET(update.ShippedAt).
		WHERE(Orders.OrderID.EQ(jet.UUID(update.OrderId))).RETURNING(Orders.AllColumns)
	} else if !update.CompletedAt.IsZero() {
		updateStmt = Orders.UPDATE(Orders.CompletedAt).
		SET(update.CompletedAt).
		WHERE(Orders.OrderID.EQ(jet.UUID(update.OrderId))).RETURNING(Orders.AllColumns)
	}
	fmt.Printf("query after Where: %s", updateStmt.DebugSql())
	res := model.Orders{}
	err := updateStmt.QueryContext(ctx, p.Client, &res)
	if err != nil {
		fmt.Errorf("Failed when updating row %w", err)
		return model.Orders{}, err
	}
	return res, nil
}

	

// List works
func (p *OrderPsqlDao) List(ctx context.Context, id uuid.UUID) ([]struct{model.Orders}, error) {
	var res []struct{
		model.Orders
	}
	getStmt := Orders.SELECT(Orders.AllColumns).FROM(Orders).WHERE(Orders.CustomerID.EQ(jet.UUID(id)))
	
	err := getStmt.QueryContext(ctx, p.Client, &res)
	if err != nil {
		fmt.Errorf("Failed to Execute List Query %w", err)
		return nil, err
	}
	
	return res, nil
}
