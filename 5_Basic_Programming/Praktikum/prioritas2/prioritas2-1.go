package main

import "fmt"

func main() {
	var jumlah int

	fmt.Printf("Input:")
	fmt.Scan(&jumlah)


	for i := 1; i <= jumlah ; i++ {
		//spasi
		for j := 1; j <= jumlah - i ; j++ {
			fmt.Printf(" ")
		}
		for x := 1; x <= i ; x++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}