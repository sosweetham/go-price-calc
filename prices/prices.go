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

func (job *TaxIncludedPriceJob) GetInputPrices() {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines, 2)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.GetInputPrices()
	result  := make(map[string]string)
	for _, price := range job.InputPrices {
		val := fmt.Sprintf("%.2f", price * (1 + job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = val
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		InputPrices: []float64{10,20,30},
		TaxRate: taxRate,
	}
}