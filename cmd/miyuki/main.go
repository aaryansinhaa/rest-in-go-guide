package main

import (
	"fmt"

	"github.com/aaryansinhaa/miyuki/pkg/config"
)

func main() {
	//Load the configuration
	cfg := config.MustLoadConfig()
	fmt.Printf("Loaded configuration: %+v\n", cfg)
	//Setup the database connection

	//setup routes


	//Start the HTTP server
	fmt.Println("Miyuki server is running...")

	// Additional setup and logic can be added here

}
