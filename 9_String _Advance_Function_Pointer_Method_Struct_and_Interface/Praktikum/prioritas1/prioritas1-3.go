package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	
	var result string

	if len(a) < len(b) {
		a , b = b , a
	}

	cek := strings.Contains(a,b)

	if cek {
		result = b
	}else{
		result = ""
	}

	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI")) // AKA
	fmt.Println(Compare("KANGOORO", "KANG")) // KANG
	fmt.Println(Compare("KI", "KIJANG")) // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA")) // ILA
}