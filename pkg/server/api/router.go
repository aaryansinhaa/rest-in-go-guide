package api

import (
	"net/http"

	"github.com/aaryansinhaa/miyuki/pkg/server/api/handlers"
	"github.com/aaryansinhaa/miyuki/pkg/storage"
)

func SetupRouter(storage storage.Storage) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Miyuki!"))
	})

	// Define different routes:
	//POST /api/block/new
	router.HandleFunc("POST /api/block/new", handlers.CreateBlock(storage))

	//GET /api/block/{id}
	router.HandleFunc("GET /api/block/{id}", handlers.GetBlock(storage))

	return router
}
