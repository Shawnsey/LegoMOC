package application

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/shawnsey/LegoMOC/LegoStore/database"
)

type App struct {
	router http.Handler
	DB     *sql.DB
}

func New() *App {

	app := &App{
		router: loadRoutes,
	}

	return app
}

func (app *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: app.router,
	}
	var err error

	app.DB, err = database.InitDB(os.Getenv("POSTGRES_URL"))
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func (app *App) CloseDB() error {
	if err := app.DB.Close(); err != nil {
		return fmt.Errorf("could not close db: %v", err)
	}
	fmt.Println("Database connection closed")
	return nil
}
