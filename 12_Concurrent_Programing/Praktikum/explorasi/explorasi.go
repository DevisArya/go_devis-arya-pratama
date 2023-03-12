package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)
const endPoint = "https://fakestoreapi.com/products"

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func getProduct(wg *sync.WaitGroup, ch chan Product) {

	response, err := http.Get(endPoint)

	if err != nil {
		log.Println(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
	}

	var responseObj []Product
	json.Unmarshal([]byte(responseData), &responseObj)

	for _, val := range responseObj{
		ch <- val
	}
	wg.Done()
}

func main() {
	
	ch := make(chan Product)

	var wg sync.WaitGroup
	fmt.Println("Product data")

	wg.Add(1)
	go getProduct(&wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch{
		fmt.Println("===")
		fmt.Println("title:", val.Title)
		fmt.Println("price:", val.Price)
		fmt.Println("category:", val.Category)
		fmt.Println("===")
	}
}