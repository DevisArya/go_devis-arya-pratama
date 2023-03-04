package main

import "fmt"

func playingDomino(cards [][]int, deck []int)interface{}  {


	sliceCard := []int{}
	
	for i := 0; i < len(cards) ; i++ {
		sliceCard = append(sliceCard, cards[i][0])
		sliceCard = append(sliceCard, cards[i][1])
	}

	for i, val := range sliceCard{
		if val == deck[0] || val == deck[1] {
			if i % 2 == 0 {
				return []int{sliceCard[i], sliceCard[i+1]}
			}
			return []int{sliceCard[i-1], sliceCard[i]}
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