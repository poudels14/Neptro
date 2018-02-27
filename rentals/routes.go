package main

import (
	"github.com/poudels14/Neptro/brbn"
	"github.com/poudels14/Neptro/rentals/controllers"
)

var Routes = []brbn.Route{
	brbn.GET("/rentals", controllers.View),
	brbn.POST("/rentals/:id", controllers.Rental),
}
