package main

import (
	"fmt"
	"math"
)

func primeNumber(number float64) bool {

	for i := 3; i <= int(math.Sqrt(number))+1; i++ {
		if int(number) % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var hasil bool
	hasil = primeNumber(100) 

	if hasil == true {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}
