package main

import (
	"fmt"

	"github.com/poudels14/bourbon/server"
)

func main() {
	s := server.Server()
	s.AddRoutes(Routes())
	s.Start()
	fmt.Println("Server starting")
}
