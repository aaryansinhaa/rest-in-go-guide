package main

import (
	"fmt"
	"net/http"

	"github.com/aaryansinhaa/miyuki/pkg/config"
)

func main() {
	//Load the configuration
	cfg := config.MustLoadConfig()
	fmt.Printf("Loaded configuration: %+v\n", cfg)
	//Setup the database connection

	//setup routes
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Miyuki! Running in %s environment", cfg.Env)
	})

	server := http.Server{
		Addr:    cfg.HTTPServerConfig.Address,
		Handler: router,
	}

	//Start the HTTP server
	fmt.Printf("Miyuki is running, open http://%s on your browser\n", cfg.HTTPServerConfig.Address)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	// Additional setup and logic can be added here

}
