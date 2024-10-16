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

type CreationsDao interface {
	Insert(ctx context.Context, model model.Creations) error
	Delete(ctx context.Context, id int32) error
	Update(ctx context.Context, id int32) (model.Creations, error)
}

type CreationsPsqlDao struct {
	Client *sql.DB
}

func NewCreationsPsqlDao(db *sql.DB) *CreationsPsqlDao {
	return &CreationsPsqlDao{Client: db}
}

func (p *CreationsPsqlDao) Insert(ctx context.Context, model model.Creations) error {
	model.ID = rand.Int32()

	insertStmt := table.Order.INSERT(table.Creations.AllColumns).MODEL(model)

	_, err := insertStmt.ExecContext(ctx, p.Client)
	if err != nil {
		fmt.Errorf("Failed to insert data: %w", err)
		return err
	}
	return nil
}

func (p *CreationsPsqlDao) Delete(ctx context.Context, id int32) error {
	deleteStmt := table.Order.DELETE().WHERE(table.Creations.ID.EQ(Int(int64(id))))

	_, err := deleteStmt.ExecContext(ctx, p.Client)
	if err != nil {
		fmt.Errorf("Entry failed to be deleted %w", err)
		return err
	}
	return nil
}

func (p *CreationsPsqlDao) Update(ctx context.Context, id int32) (model.Creations, error) {

	return model.Creations{}, nil 
}