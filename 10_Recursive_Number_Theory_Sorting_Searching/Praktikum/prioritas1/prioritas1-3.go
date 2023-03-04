package main

import (
	"fmt"
	"math"
)


func primeNumber(number int) bool {

	if number < 2{
		return false
	}
	if number == 2 {
		return true
	}

	for i := 2; i <= int(math.Sqrt(float64(number)))+1; i++ {
		if int(number) % i == 0 {
			return false
		}
	}

	return true
}

func primeX(number int)int  {
	var result int

	for i := 2; i > 0; i++ {
		if number == 0 {
			break
		}

		if primeNumber(i){
			number--
		}
		result = i
	}

	return result
}

func main() {
   	fmt.Println(primeX(1))         //2
   	fmt.Println(primeX(5))         //11
   	fmt.Println(primeX(8))         //19
   	fmt.Println(primeX(9))         //23
   	fmt.Println(primeX(10))         //29
}
