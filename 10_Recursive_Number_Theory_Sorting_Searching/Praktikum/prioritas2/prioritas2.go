package main

import "fmt"

func playingDomino(cards [][]int, deck []int)interface{}  {
	
	for i, val := range cards{

		if val[0] == deck[0] || val[0] == deck[1] || val[1] == deck[0] || val[1] == deck[1]  {
			return cards[i]
		}
		
	}
	return "tutup kartu"
}

func main()  {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3,4}, []int{2,1}, []int{3,3}}, []int{4,3}))
	// //[3, 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3,3}, []int{3,4}, []int{2,1}}, []int{3,6}))
	// [6 5]
	fmt.Println(playingDomino([][]int{[]int{6,6}, []int{2,4}, []int{3,6}}, []int{5, 1}))
	// "tutup kartu"
}