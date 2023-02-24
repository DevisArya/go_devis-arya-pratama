package main

import (
	"fmt"
)

func main() {
	//calculate pow of integers inputs
	a := 2
	b := 17
	result := calPowIntRecur(a, b)
	fmt.Println(result)
}
var counter = 0
// Assumption that n >= 0
func calPowIntRecur(x, n int) int {
	// fmt.Println("test")
	fmt.Println(n)
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
       counter = counter + 1 

	   // recursion here
	y := calPowIntRecur(x, n/2)
	// fmt.Println("========")
	fmt.Println("ini n", n)
	
	
	if n%2 == 0 {
		
		return y * y
	}
	return x * y * y
}
