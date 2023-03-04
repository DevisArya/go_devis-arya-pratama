package main

import (
	"fmt"
)

func pow(x, n int) int {

	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}
	
	var k = pow(x, n/2)
	
	if n % 2 == 0 {
		return k * k
	}else{
		return x * k * k
	}
}

func main() {
   fmt.Println(pow(2, 3))  // 8
   fmt.Println(pow(5, 3))  // 125
   fmt.Println(pow(10, 2)) // 100
   fmt.Println(pow(2, 5))  // 32
   fmt.Println(pow(7, 3))  // 343
}

