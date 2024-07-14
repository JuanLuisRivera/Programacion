package main

import (
	cmdmanager "example.com/m/CMDManager"
	prices "example.com/m/Prices"
)

func main() {
	taxRate := []float64{0.0, 0.05, 0.1}

	for _, tax := range taxRate {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmd, tax)
		priceJob.Process()
	}
}
