package main

import (
	"price-calculator/price"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := price.NewTaxIncludedPriceJob(taxRate)
		priceJob.Processed()
	}
}
