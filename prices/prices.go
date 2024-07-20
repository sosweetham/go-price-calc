package prices

import (
	"fmt"

	"kodski.com/price-calculator/conversion"
	"kodski.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager iomanager.IOManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) GetInputPrices() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines, 2)

	if err != nil {
		return err 
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.GetInputPrices()

	if err != nil {
		return err
	}

	result  := make(map[string]string)
	for _, price := range job.InputPrices {
		val := fmt.Sprintf("%.2f", price * (1 + job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = val
	}
	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		InputPrices: []float64{10,20,30},
		TaxRate: taxRate,
	}
}