package main

import (
	"github.com/poudels14/Neptro/brbn"
	"github.com/poudels14/Neptro/brbn/middleware"
	"github.com/poudels14/Neptro/users/controllers"
)

func main() {
	server := brbn.New("0.0.0.0", "3000")

	// adding middleware
	server.Chain(middleware.Logger)

	// adding routes
	server.GET("/users", controllers.GetUsers)
	server.GET("/users/:id", controllers.GetUser)
	server.POST("/users/new", controllers.CreateUser)

	// starting server...
	server.Start()
}
