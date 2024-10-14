package daos

import (
	"database/sql"
)

type CreationsDao struct {
	Client *sql.DB
}

func NewCreationDao(db *sql.DB) *CreationsDao {
	return &CreationsDao{Client: db}
}