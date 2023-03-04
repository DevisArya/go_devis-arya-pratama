package main

import "fmt"

func caesar(offset int, input string) string {

	offset = offset % 26
	data := []byte(input)
	
	for i, val := range data{

		if val + byte(offset) > 122{
			data[i] = 96 + ((byte(offset))-(122-val))
		}else{
			data[i] = val + byte(offset)
		}
	}
	
	result := string(data)

	
	return result
	
	}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // cnvc
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi 
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza 
  fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl 
}