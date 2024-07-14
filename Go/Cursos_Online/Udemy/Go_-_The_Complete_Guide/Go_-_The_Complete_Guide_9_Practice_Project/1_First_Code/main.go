package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRate := []float64{0, 0.10, 0.15}

	result := make(map[float64][]float64)

	for _, tax := range taxRate {
		taxedPrices := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxedPrices[priceIndex] = price * (1 + tax)
		}
		result[tax] = taxedPrices
	}

	fmt.Println(result)
}
