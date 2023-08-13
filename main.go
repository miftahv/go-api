package main

import (
	"go-api/db"
	"go-api/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":6361"))
}
