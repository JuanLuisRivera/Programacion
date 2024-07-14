package main

import (
	prices "example.com/m/Prices"
)

func main() {
	taxRate := []float64{0.0, 0.05, 0.1}

	for _, tax := range taxRate {
		priceJob := prices.NewTaxIncludedPriceJob(tax)
		priceJob.Process()
	}
}
