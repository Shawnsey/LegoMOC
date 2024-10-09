package database

import (
	"database/sql"
	"fmt"
)

func InitDB(connectionString string) (*sql.DB, error) {
	var err error
	DB, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %v", err)
	}

	// Check if the database is actually reachable
	err = DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not connect to db: %v", err)
	}

	fmt.Println("Database connected successfully")
	return DB, nil
}
