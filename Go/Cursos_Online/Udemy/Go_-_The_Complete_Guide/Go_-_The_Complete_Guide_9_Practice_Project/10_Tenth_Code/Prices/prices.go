package prices

import (
	"fmt"

	conversion "example.com/m/Conversion"
	iomanager "example.com/m/IOManager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadPrices()

	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(*job)
}

func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   io,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
