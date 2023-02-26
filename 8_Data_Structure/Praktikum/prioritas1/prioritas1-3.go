package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {
	
	pecah := strings.Split(angka,"")
	sliced := make([]int, len(pecah))
	maps := make(map[int]bool)
	del := []int{}
	result := []int{}

	for i := range sliced{
		sliced[i], _ = strconv.Atoi(pecah[i])
	}
	
	for _, val := range sliced{

		if value := maps[val] ; !value {
			maps[val] = true
		}else{
			del = append(del, val)	
		}
	}
	for _, v := range del{
		delete(maps, v)
	}
	for key := range maps{
		result = append(result, key)
	}

	return result
}

func main()  {
	fmt.Println(munculSekali("1234123")) 	   // [4]
	fmt.Println(munculSekali("76523752"))      // [6, 3]
	fmt.Println(munculSekali("12345"))		   // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455"))    // []
	fmt.Println(munculSekali("0872504"))       // [8 7 2 5 4]
}
