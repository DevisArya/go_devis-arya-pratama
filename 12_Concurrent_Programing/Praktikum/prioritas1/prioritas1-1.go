package main

import (
	"fmt"
	"time"
)

func printNumber(number int) {

	for i := 1; i <= number; i++ {
		fmt.Println(i)
		time.Sleep(3 * time.Second)
	}
}
func main() {
	num := 5
	go printNumber(5)
	time.Sleep(time.Duration(num)* 3 * time.Second)
}