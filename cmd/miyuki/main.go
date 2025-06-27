package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/aaryansinhaa/miyuki/pkg/config"
	"github.com/aaryansinhaa/miyuki/pkg/server"
	"github.com/aaryansinhaa/miyuki/pkg/storage/sqlite"
)

func main() {
	//Load the configuration
	cfg := config.MustLoadConfig()
	fmt.Printf("Loaded configuration: %+v\n", cfg)
	//Setup the database connection
	storage, err := sqlite.LoadSQLiteStorage(cfg)
	if err != nil {
		log.Fatal("Failed to initialize storage:", err)
	}
	slog.Info("Storage initialized successfully", "storagePath", cfg.StoragePath)
	//Start the local server

	server.LocalServer(cfg, storage)

	// Additional setup and logic can be added here

}
