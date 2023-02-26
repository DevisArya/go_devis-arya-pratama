package main

import "fmt"


func pairSum(arr []int, target int) [] int{
	result := []int{}

	for i := 0; i < len(arr); i++ {
		for k := i+1; k < len(arr); k++ {
			if arr[i] + arr[k] == target {
				result = append(result, i)
				result = append(result, k)
			}
		}
	}
	return result
}

func main(){
	fmt.Println(pairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(pairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]
	fmt.Println(pairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]
	fmt.Println(pairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]
	fmt.Println(pairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]
}