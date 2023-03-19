package main

import (
	"fmt"
	"time"
)

func main() {
	number := 10

	ch := make(chan int)
	
	go func(num int) {
		for i := 1; i <= num; i++ {
			ch <- i*3
		}
	}(number)

	go func() {
		for{
			receiveCh := <- ch
			fmt.Println(receiveCh)
		}
	}()
	
	<-time.After(time.Second)
}