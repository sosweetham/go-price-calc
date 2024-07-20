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
Tax Rates are: 0%, 7%, 10%, 15%
Press Enter after selecting an option (1 or 2)
Please select an option:
1. Enter your prices manually
2. Read prices from a file
		`)
		
		var option int
		fmt.Scan(&option)

		var manager iomanager.IOManager

		if option == 1 {
			manager = cmdmanager.NewCMDManager()
		} else if option == 2 {
			manager = filemanager.NewFileManager("input.txt", "output.txt")
		} else {
			fmt.Println("Invalid option")
			continue
		}

		for _, tax := range taxes {
			priceJob := prices.NewTaxIncludedPriceJob(manager, tax)
			priceJob.Process()
		}
	}
}