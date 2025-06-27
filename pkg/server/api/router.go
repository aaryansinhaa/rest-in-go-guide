package api

import (
	"net/http"

	"github.com/aaryansinhaa/miyuki/pkg/server/api/handlers"
)

func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Miyuki!"))
	})

	// Define different routes:
	router.HandleFunc("/api/block", handlers.CreateBlock())

	return router
}
