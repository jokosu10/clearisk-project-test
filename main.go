package main

import (
	"project-alta-store/routes"
)

func main() {
	e := routes.Start()

	e.Logger.Fatal(e.Start(":8000"))
}
