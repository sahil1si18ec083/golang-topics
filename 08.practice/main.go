package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {

	// pricess := []float64{10, 20, 30}

	taxesRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxesRates {
		f := filemanager.Filemanager{InputFileName: "prices.txt", OutputFileName: fmt.Sprintf("%.2f data.json", taxRate*100)}
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, f)

		priceJob.Process()

	}

}
