package main

import "fmt"

type Car struct {
	fuelln float64
	tipe   string
}

func (c Car) GetDistance() float64 {
	var dist float64
	if c.tipe == "Civic" {
		dist = 4.5
	} else if c.tipe == "innova" {
		dist = 6
	} else if c.tipe == "Avanza" {
		dist = 10
	} else {
		dist = 4
	}
	return c.fuelln * dist
}

func main() {
	car := Car{17.5, "innova"}
	fmt.Println("Jarak yang bisa ditempuh adalah", car.GetDistance(), "Mill")
}