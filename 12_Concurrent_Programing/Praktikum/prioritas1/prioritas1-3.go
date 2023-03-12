package main

import (
	"fmt"
	"sync"
)

func sendCh(wg *sync.WaitGroup, ch chan int, num int)  {
	for i := 1; i <= 10; i++ {
		ch <- i * 3
	}
	wg.Done()
}

func main() {
	number := 10
	var wg sync.WaitGroup

	ch := make(chan int, 2)
	
	wg.Add(1)
	go sendCh(&wg, ch, number)
	
	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch{
		fmt.Println(val) 
	}
	
}