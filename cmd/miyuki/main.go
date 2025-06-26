package main

import (
	"fmt"

	"github.com/aaryansinhaa/miyuki/pkg/config"
	"github.com/aaryansinhaa/miyuki/pkg/server"
)

func main() {
	//Load the configuration
	cfg := config.MustLoadConfig()
	fmt.Printf("Loaded configuration: %+v\n", cfg)
	//Setup the database connection

	//Start the local server

	server.LocalServer(cfg)

	// Additional setup and logic can be added here

}
