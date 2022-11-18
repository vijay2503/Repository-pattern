package main

import (
	"go-postgres/repository"
	"go-postgres/router"
)

var userRepo repository.UserRepo

func main() {
	e := router.Routing()
	e.Logger.Fatal(e.Start(":8000"))
}
