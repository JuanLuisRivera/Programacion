package prices

import (
	"fmt"

	conversion "example.com/m/Conversion"
	filemanager "example.com/m/FileManager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(*job)
}

func NewTaxIncludedPriceJob(file filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   file,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
