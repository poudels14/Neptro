package main

import (
	"github.com/poudels14/Neptro/brbn"
)

func main() {
	server := brbn.Server{Address: "0.0.0.0", Port: "5555"}

	// adding routes...
	// b.Get("/rental/:id", "get_rental", nil)
	// b.Post("/rental", "create_rental", nil)

	server.Chain(
	// Put middlewares in order of execution
	)
	server.AddRoutes(Routes)

	// starting server...
	server.Start()
}
