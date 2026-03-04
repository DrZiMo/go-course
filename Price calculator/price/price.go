package price

import (
	"fmt"
	"price-calculator/convertions"
	"price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate          float64           `json:"tax_rate"`
	InputPrices      []float64         `json:"input_price"`
	TaxIncludedPrice map[string]string `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("data/prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := convertions.StringToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Processed() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrice = result
	filemanager.WriteJSON(fmt.Sprintf("data/result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
