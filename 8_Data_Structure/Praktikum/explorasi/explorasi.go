package main

import (
	"fmt"
	"math"
)
func selisihDiagonal(panjang int)float64 {
	diagonal1 := 0
	diagonal2 := 0

	matrix := make([][]int, panjang)

	for i := 0; i < panjang; i++{
		
		matrix[i] = make([]int,panjang)
		fmt.Println("Masukan Baris ke",i+1,", contoh jika panjangnya 3 = 1 2 3")
		
		for k := 0; k < panjang ;k++{
			
			fmt.Scan(&matrix[i][k])
			if i == k {
				diagonal1 += matrix[i][k]
			}
			if i+k == panjang-1 {
				diagonal2 += matrix[i][k]
			}
		}
	}

	selisih := math.Abs(float64(diagonal1 - diagonal2))
	
	return selisih
}

func main() {
	var panjang int
	fmt.Printf("Masukan Panjang Matrix:")
	fmt.Scan(&panjang)
	
	fmt.Println("Perbedaan mutlak adalah", selisihDiagonal(panjang))
}