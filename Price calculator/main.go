package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/price"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("data/prices.txt", fmt.Sprintf("data/result_%.0f.json", taxRate*100))
		priceJob := price.NewTaxIncludedPriceJob(*fm, taxRate)
		go priceJob.Processed(doneChans[index])
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}
