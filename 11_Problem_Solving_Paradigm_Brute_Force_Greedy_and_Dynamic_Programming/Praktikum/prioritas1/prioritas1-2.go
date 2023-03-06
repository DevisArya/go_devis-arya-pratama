package main

import "fmt"

func number(number int) [][]int  {

	result := make([][]int, number)

	for i := 0; i < number; i ++ {
        result[i] = make([]int, i + 1)
        result[i][0], result[i][i] = 1, 1
        for k := 1; k < i; k ++ {
            result[i][k] = result[i - 1][k] + result[i - 1][k - 1]
        }
    }
    return result
}

func main()  {
	fmt.Println(number(5))
}