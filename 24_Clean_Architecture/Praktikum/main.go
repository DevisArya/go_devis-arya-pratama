package main

import (
	"praktikum/config"
	"praktikum/route"
)

func main() {
	config.InitDB()
	router := route.StartRoute()

	router.Start(":8000")
}
