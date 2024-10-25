package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"github.com/joho/godotenv"

	"github.com/shawnsey/LegoMOC/LegoStore/application"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
	defer app.CloseDB()

}
