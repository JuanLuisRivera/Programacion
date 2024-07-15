package main

import (
	"fmt"

	filemanager "example.com/m/FileManager"
	prices "example.com/m/Prices"
)

func main() {
	taxRate := []float64{0.0, 0.05, 0.1}
	doneChan := make([]chan bool, len(taxRate))
	errorChan := make([]chan error, len(taxRate))

	for index, tax := range taxRate {
		doneChan[index] = make(chan bool)
		errorChan[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		//cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, tax)
		go priceJob.Process(doneChan[index], errorChan[index])
	}

	for index, _ := range taxRate { // We use the taxRate slice as it is the one that controls the number of channels created
		select { // We use the "select" function with channels to specify that we will use a "case" structure, where we will wait for two different channels, and the one that emits a value first is the one that counts and it proceeds with the next iteration of channels
		case err := <-errorChan[index]: // If the channel has an error it will print the error
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChan[index]: // If the channel completed the job it will print the statement "Done"
			fmt.Println("Done")
		}
	}
}
