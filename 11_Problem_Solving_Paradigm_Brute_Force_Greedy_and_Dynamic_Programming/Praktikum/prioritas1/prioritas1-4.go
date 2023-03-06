package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) interface{} {

	if a < 3 {
		return "no solution"
	}

	for x := 1; x <= a/3; x++ {
		for y := 1; y <= b; y++ {
			for z := 1; z <= c/3; z++ {
				if x + y + z == a {
					if x * y * z == b {
						if math.Pow(float64(x),2)+math.Pow(float64(y),2)+math.Pow(float64(z),2) == float64(c) {
							return []int {x,y,z}
						}
					}
				}
			}
		}
	}

	return "no solution"
}

func main() {
	fmt.Println(SimpleEquations(1, 2, 3)) // no solution
	fmt.Println(SimpleEquations(6, 6, 14)) // 1 2 3
}
