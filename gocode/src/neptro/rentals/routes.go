package main

import (
	"neptro/rentals/controllers"

	"github.com/poudels14/bourbon/router"
)

func Routes() []router.RouteStruct {
	return []router.RouteStruct{
		router.Route("GET", "/rentals", controllers.View),
		router.Route("GET", "/rentals/{id}", controllers.Rental),
	}
}
