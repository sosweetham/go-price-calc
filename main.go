package main

import (
	"fmt"

	"kodski.com/price-calculator/filemanager"
	"kodski.com/price-calculator/prices"
)

func main() {
	taxes := []float64{0, 0.07, 0.1, 0.15}
	for _, tax := range taxes {
		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("output_%.0f.json", tax*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, tax)
		priceJob.Process()
	}
}