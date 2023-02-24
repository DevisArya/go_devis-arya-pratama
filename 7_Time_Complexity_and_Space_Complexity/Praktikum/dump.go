package main

import "fmt"

func checkPrime(number int) bool {
	if number == 1 {
		return false
	}
	for i := 2; i <= number / 2; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var hasil bool
	hasil = checkPrime(107) // false

	if hasil == true {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}