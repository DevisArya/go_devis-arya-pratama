package main

import (
	"fmt"
	"math"
)

func primeNumber(number float64) bool {

	if number < 2{
		return false
	}
	if number == 2 {
		return true
	}

	for i := 3; i <= int(math.Sqrt(number))+1; i++ {
		if int(number) % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var hasil bool

	hasil = primeNumber(1000000007) // true
   	// hasil = primeNumber(13)         // true
   	// hasil = primeNumber(17)         // true
   	// hasil = primeNumber(20)         // false
   	// hasil = primeNumber(35)         // false

	if hasil == true {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}
