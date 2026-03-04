package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/price"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("data/prices.txt", fmt.Sprintf("data/result_%.0f.json", taxRate*100))
		priceJob := price.NewTaxIncludedPriceJob(*fm, taxRate)
		priceJob.Processed()
	}
}
