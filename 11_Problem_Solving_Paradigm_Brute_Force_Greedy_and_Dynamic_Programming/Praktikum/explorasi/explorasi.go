package main

import "fmt"


func numToRomawi(number int)(result string) {

	angka:= []int{1,4,5,9,10,40,50,90,100,400,500,900,1000}
	romawi:= []string{"I","IV","V","IX","X","XL","L","XC","C","CD","D","CM","M"}

	for i := len(romawi)-1; i >= 0; i-- {
		
		for number >= angka[i] {
			number -= angka[i]
			result += romawi[i]
		}
	}
	return result
}

func main()  {
	fmt.Println(numToRomawi(4))
	fmt.Println(numToRomawi(9))
	fmt.Println(numToRomawi(23))
	fmt.Println(numToRomawi(2021))
	fmt.Println(numToRomawi(1646))
}