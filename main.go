package main

import (
	db "echo-framework/db"
	routes "echo-framework/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))
}
