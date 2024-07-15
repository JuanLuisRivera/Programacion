package main

import (
	"fmt"

	filemanager "example.com/m/FileManager"
	prices "example.com/m/Prices"
)

func main() {
	taxRate := []float64{0.0, 0.05, 0.1}
	doneChan := make([]chan bool, len(taxRate))

	for index, tax := range taxRate {
		doneChan[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		//cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, tax)
		go priceJob.Process(doneChan[index])

		//if err != nil {
		//	fmt.Println("could not process job")
		//	fmt.Println(err)
		//}
	}

	for _, done := range doneChan {
		<-done
	}
}
