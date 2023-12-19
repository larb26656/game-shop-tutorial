package main

import (
	"context"
	"log"
	"os"

	"github.com/larb26656/game-shop-tutorial/config"
	"github.com/larb26656/game-shop-tutorial/pkg/database"
	"github.com/larb26656/game-shop-tutorial/server"
)

func main() {
	ctx := context.Background()
	// Initialize config

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		} 

		return os.Args[1]
	}())

	// Connect to database
	db := database.DbCon(ctx, &cfg)
	defer db.Disconnect(ctx)

	// Start server
	server.Start(ctx, &cfg, db)
}