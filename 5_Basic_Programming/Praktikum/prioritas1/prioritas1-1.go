package main

import "fmt"

func luasTrapesium(sisiA float64, sisiB float64, tinggi float64) float64 {
	return ((sisiA + sisiB) * tinggi) / 2
}

func main() {
	var sisiA, sisiB, tinggi float64

	fmt.Printf("Masukan panjang sisi A:")
	fmt.Scan(&sisiA)
	fmt.Printf("Masukan panjang sisi B:")
	fmt.Scan(&sisiB)
	fmt.Printf("Masukan tinggi:")
	fmt.Scan(&tinggi)
	fmt.Println(luasTrapesium(sisiA,sisiB,tinggi))
}