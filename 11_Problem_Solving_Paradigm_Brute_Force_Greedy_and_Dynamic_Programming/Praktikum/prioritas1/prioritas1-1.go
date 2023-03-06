package main

import (
	"fmt"
	"strconv"
)

func cetakBiner(number int) (result []string) {
	
	for i := 0; i <= number; i++ {
		result = append(result, strconv.FormatInt(int64(i), 2))
	}
	return result
}

func main() {
	fmt.Println(cetakBiner(2))
	fmt.Println(cetakBiner(3))
}