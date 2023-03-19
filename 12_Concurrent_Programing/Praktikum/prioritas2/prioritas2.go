package main

import (
	"fmt"
	"sync"
)

type Letters struct {
	m sync.Mutex
	result map[string]int
}

func (l *Letters) countLetters(wg *sync.WaitGroup,s string) {

	l.m.Lock()
	defer l.m.Unlock()

	_, cek := l.result[s]

	if cek {
		l.result[s] += 1
	} else {
		l.result[s] = 1
	}
	wg.Done()
}

func main() {
	var sentence string = "qetsafdhamjsasdaghihuiwqduhhsaoi"

	var wg sync.WaitGroup

	lt := Letters{
		result: map[string]int{},
	}

	for _, val := range sentence {
		wg.Add(1)
		go lt.countLetters(&wg, string(val))
	}
	wg.Wait()

	fmt.Println(lt.result)

}