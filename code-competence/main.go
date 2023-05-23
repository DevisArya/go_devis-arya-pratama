package main

import (
	"code-competence/config"
	"code-competence/routes"
)

func main() {
	config.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
