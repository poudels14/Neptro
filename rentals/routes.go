package main

import (
	"github.com/poudels14/Neptro/brbn"
	"github.com/poudels14/Neptro/rentals/controllers"
)

func AllRoutes() []brbn.Route {
	return []brbn.Route{
		brbn.GET("/rentals", controllers.View),
		brbn.POST("/rentals/:id", controllers.Rental),
	}
}
