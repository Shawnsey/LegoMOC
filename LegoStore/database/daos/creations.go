package daos

import (
	"context"
	"database/sql"
	"fmt"

	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/model"
	. "github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/table"
)

type CreationsDao interface {
	Insert(ctx context.Context, model model.Creations) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID) (model.Creations, error)
	List(ctx context.Context, id uuid.UUID) ([]model.Creations, error)
}

type CreationsPsqlDao struct {
	Client *sql.DB
}

func NewCreationsPsqlDao(db *sql.DB) *CreationsPsqlDao {
	return &CreationsPsqlDao{Client: db}
}

func (p *CreationsPsqlDao) Insert(ctx context.Context, model model.Creations) error {
	creation_id, err := uuid.NewUUID()
	if err != nil {
		fmt.Errorf("Failed to create uuid",err)
		return err
	}
	model.ID= creation_id

	insertStmt := Creations.INSERT(Creations.AllColumns).MODEL(model)

	_, err = insertStmt.ExecContext(ctx, p.Client)
	if err != nil {
		fmt.Errorf("Failed to insert data: %w", err)
		return err
	}
	return nil
}

func (p *CreationsPsqlDao) Delete(ctx context.Context, id uuid.UUID) error {
	deleteStmt := Creations.DELETE().WHERE(Creations.ID.IN(jet.UUID(id)))

	_, err := deleteStmt.ExecContext(ctx, p.Client)
	if err != nil {
		fmt.Errorf("Entry failed to be deleted %w", err)
		return err
	}
	return nil
}

func (p *CreationsPsqlDao) Update(ctx context.Context, id uuid.UUID) (model.Creations, error) {

	return model.Creations{}, nil 
}

func (p *CreationsPsqlDao) List(ctx context.Context, id uuid.UUID) ([]model.Creations, error) {
	return []model.Creations{}, nil
}