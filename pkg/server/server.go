package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aaryansinhaa/miyuki/pkg/config"
	"github.com/aaryansinhaa/miyuki/pkg/server/api"
	"github.com/aaryansinhaa/miyuki/pkg/storage"
)

func LocalServer(cfg *config.Config, storage storage.Storage) {

	// Setup routes
	router := api.SetupRouter(storage)

	server := http.Server{
		Addr:    cfg.HTTPServerConfig.Address,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {

		// Start the HTTP server
		fmt.Printf("Miyuki is running, open http://%s on your browser\n", cfg.HTTPServerConfig.Address)

		err := server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}

	}()

	<-done

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error shutting down server", "error", err)
	}
	slog.Info("Server gracefully stopped")
}
