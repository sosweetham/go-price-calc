package prices

import (
	"fmt"
	"strconv"

	"kodski.com/price-calculator/conversion"
	"kodski.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) GetInputPrices() {
	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices := make([]float64, len(lines))

	for _, line := range lines {
		initialVal, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		val, err := conversion.LimitFloat(initialVal, 2)
		if err != nil {
			fmt.Println(err)
			return
		}
		prices = append(prices, val)
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.GetInputPrices()
	result  := make(map[string]float64)
	for _, price := range job.InputPrices {
		val, err := strconv.ParseFloat(fmt.Sprintf("%.2f", price * (1 + job.TaxRate)), 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		result[fmt.Sprintf("%.2f", price)] = val
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10,20,30},
		TaxRate: taxRate,
	}
}