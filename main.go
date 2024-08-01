package main

import (
	"github.com/JulisTafita/go-savannahTech/internal/db"
	"github.com/JulisTafita/go-savannahTech/pkg/server"
	"github.com/JulisTafita/go-savannahTech/pkg/worker"
)

func main() {
	// Ensure that the database connection is closed when the main function exits.
	defer db.CloseConnexion()

	// Start the web server in a separate goroutine.
	go func() {
		// Call the Serve function from the server package to start the web server.
		err := server.Serve()
		if err != nil {
			// If there's an error starting the server, panic with the error message.
			panic(err)
		}
	}()

	// Run the worker job, which handles periodic tasks.
	err := worker.Run()
	if err != nil {
		// If there's an error running the worker, panic with the error message.
		panic(err)
	}
}
