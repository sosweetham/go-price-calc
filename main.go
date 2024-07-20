package main

import (
	"fmt"

	"kodski.com/price-calculator/cmdmanager"
	"kodski.com/price-calculator/filemanager"
	"kodski.com/price-calculator/iomanager"
	"kodski.com/price-calculator/prices"
)

func main() {
	for {
		taxes := []float64{0, 0.07, 0.1, 0.15}
	
		fmt.Println(`
Welcome to Tax Included Price Calculator
Press Enter after selecting an option (1 or 2)
Please select an option:
1. Enter your prices manually
2. Read prices from a file
3. Exit
		`)

		type option int 

		const (
			manual option = iota + 1
			file
			exit
		)

		var selectedOption option
		fmt.Scan(&selectedOption)

		if selectedOption == exit {
			break
		}

		var iom iomanager.IOManager

		for _, tax := range taxes {

			switch selectedOption {
			case manual:
				iom = cmdmanager.NewCMDManager()
			case file:
				iom = filemanager.NewFileManager("prices.txt", fmt.Sprintf("output_%.0f.json", tax*100))
			}

			priceJob := prices.NewTaxIncludedPriceJob(iom, tax)
			err := priceJob.Process()

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}