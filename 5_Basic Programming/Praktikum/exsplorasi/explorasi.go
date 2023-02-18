package main

import (
	"bufio"
	"fmt"
	"os"
)

func cekPalindrome(kata string) bool {
	for i := 0; i < len(kata)/2; i++ {
		if kata[i] != kata[len(kata)-i-1]{
			return false
		}
	}
	return true
}

func main() {
	var hasil bool

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Masukan Kata :")
	scanner.Scan()
	kata := scanner.Text()

	hasil = cekPalindrome(kata)
	if hasil == true{
		fmt.Println("Palindrome")
	}else{
		fmt.Println("Bukan Palindrome")
	}	
}