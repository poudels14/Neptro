package main

import (
	"github.com/poudels14/Neptro/brbn"
	"github.com/poudels14/Neptro/rentals/controllers"
)

func main() {
	server := brbn.New("0.0.0.0", "5555")

	// adding middleware
	passthrough := func(handler brbn.Handler) brbn.Handler { return handler }
	server.Chain(passthrough).Chain(passthrough).Chain(passthrough)

	// adding routes
	server.GET("/rentals", controllers.View)
	server.POST("/rentals/:id", controllers.Rental)

	// starting server...
	server.Start()
}
