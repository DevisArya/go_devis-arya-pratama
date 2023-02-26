package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	
	arrayA = append(arrayA, arrayB...)
	maps := make(map[string]bool)
	result := []string{}

	for _, val := range arrayA{

		if value := maps[val] ; !value {

			maps[val] = true
			
			result = append(result, val)
		}
	}
	return result
}
func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}