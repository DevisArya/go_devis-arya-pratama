package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int)int  {
	
	minimum, temp := 0, 0
	for i := 1; i < len(jumps); i++ {
		
		jump1 := minimum + int(math.Abs(float64(jumps[i])-float64(jumps[i-1])))
		jump2 := math.MaxInt
		if i > 1 {
			jump2 = temp + int(math.Abs(float64(jumps[i])-float64(jumps[i-2])))
		}
		cekMin := math.Min(float64(jump1), float64(jump2))
		temp = minimum
		minimum = int(cekMin)
	}
	return minimum
}

func main()  {
	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}