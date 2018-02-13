package main

import (
	"github.com/poudels14/Neptro/brbn"
)

func main() {
	b := brbn.Default()

	// adding routes...
	b.Get("/rental/:id", "get_rental", nil)
	b.Post("/rental", "create_rental", nil)

	// starting server...
	b.Start()
}
