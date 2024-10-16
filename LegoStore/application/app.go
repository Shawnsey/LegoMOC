package application

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/shawnsey/LegoMOC/LegoStore/database"
)

type App struct {
	router http.Handler
	DB     *sql.DB
}

func New() *App {

	app := &App{}
	app.loadRoutes()

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
		return fmt.Errorf("failed to connect to db: %w", err)
	}


	defer func() {
		if err := app.CloseDB(); err != nil {
			fmt.Println("failed to close db", err)
		}
	}()

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}

func (app *App) CloseDB() error {
	if err := app.DB.Close(); err != nil {
		return fmt.Errorf("could not close db: %v", err)
	}
	fmt.Println("Database connection closed")
	return nil
}
