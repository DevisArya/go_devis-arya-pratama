package main

import "fmt"

type pair struct {
	name string
	count int
}

func mergeSort(elements []pair)[]pair  {

	var panjang = len(elements)

	if panjang == 1{
		return elements
	}

	tengah := int(panjang/2)

	kiri := make([]pair, tengah)
	kanan := make([]pair, panjang-tengah)

	for i := 0; i < panjang; i++ {
		
		if i < tengah {
			kiri[i] = elements[i]
				
		}else{
			kanan[i-tengah] = elements[i]
		}
	}

	return merge(mergeSort(kiri), mergeSort(kanan))
}

func merge(kiri, kanan []pair) (result []pair) {

	result = make([]pair, len(kiri)+len(kanan))

	i := 0

	for len(kiri) > 0 && len(kanan) > 0{

		if kiri[0].count < kanan[0].count {
			result[i] = kiri[0]
			kiri = kiri[1:]
		}else{
			result[i] = kanan[0]
			kanan = kanan[1:]
		}
		i++
	}

	for k := 0; k < len(kiri); k++ {
		result[i] = kiri[k]
		i++
	}
	for k := 0; k < len(kanan); k++ {
		result[i] = kanan[k]
		i++
	}

	return result
}

func MostAppearItem(items []string)[]pair  {

	unsortedData := make(map[string]int)

	for _, val := range items {

		_, cek := unsortedData[val]
		
		if cek{
			unsortedData[val] += 1
		}else{
			unsortedData[val] = 1
		}
	}

	result := []pair{}
	var input pair
	for key, val := range unsortedData {
		input.name = key
		input.count = val
		result = append(result, input)
	}

	result = mergeSort(result)

	return result
}

func main()  {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang -> 1 ruby -> 2 js -> 4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C-> 1 D -> 2 B-> 3 A->4 
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// footbal -> 1 basketball -> 1 tenis -> 1
}