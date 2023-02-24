package main

import (
	"fmt"
)

func pow(x, n int) int {
	var hasil int = x
	for n > 1 {
		fmt.Println(n)

		if n%2 == 0 {
			hasil *= hasil
		}else{
			hasil = x * hasil * hasil
		}
		n /= 2
	}
	return hasil
}

func main() {
   fmt.Println(pow(5, 17))  
}