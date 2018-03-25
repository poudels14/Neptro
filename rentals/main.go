package main

import (
	"github.com/poudels14/Neptro/brbn"
	"github.com/poudels14/Neptro/brbn/middleware"
	"github.com/poudels14/Neptro/rentals/controllers"
)

func main() {
	server := brbn.New("0.0.0.0", "5555")

	// adding middleware
	server.Chain(middleware.Logger)

	// adding routes
	server.GET("/rentals", controllers.GetRentals)
	server.GET("/rentals/:id", controllers.GetRental)
	server.POST("/rentals/new", controllers.CreateRental)

	// starting server...
	server.Start()
}
