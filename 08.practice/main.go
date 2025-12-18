package main

import (
	"price-calculator/prices"
)

func main() {

	// pricess := []float64{10, 20, 30}

	taxesRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxesRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)

		priceJob.Process()

	}

}
