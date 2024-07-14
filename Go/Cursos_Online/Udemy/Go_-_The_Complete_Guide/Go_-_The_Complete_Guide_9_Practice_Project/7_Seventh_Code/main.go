package main

import (
	"fmt"

	filemanager "example.com/m/FileManager"
	prices "example.com/m/Prices"
)

func main() {
	taxRate := []float64{0.0, 0.05, 0.1}

	for _, tax := range taxRate {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		priceJob := prices.NewTaxIncludedPriceJob(*fm, tax)
		priceJob.Process()
	}
}
