package main

import "fmt"

func main() {
	var angka int64

	fmt.Printf("Masukan Angka :")
	fmt.Scan(&angka)

	if angka % 2 == 0 {
		fmt.Println("Bilangan Genap")
	}else{
		fmt.Println("Bilangan Ganjil")
	}

}