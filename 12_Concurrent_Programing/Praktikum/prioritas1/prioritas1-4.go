package main

import (
	"fmt"
	"sync"
	"time"
)

type safeNumber struct{
	result int
	m sync.Mutex
}

func (i *safeNumber) SetFibo(num int)  {
	
	i.m.Lock()
	defer i.m.Unlock()

	if num < 2{
		i.result = num
	}
	
	prev, res := 0, 1

	for i:= num-1 ; 0 < i ; i--{
		prev += res
		prev, res = res, prev
	}
	
	i.result = res
}

func (i *safeNumber) GetFibo()int  {
	i.m.Lock()
	defer i.m.Unlock()
	return i.result
}

func main() {

	i := &safeNumber{}

	go func() {
		i.SetFibo(12)
	}()
	time.Sleep(time.Second)


	fmt.Println(i.GetFibo())
}