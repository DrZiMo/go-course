package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount, years float64 = 0.0, 10
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value (Adjust it for inflation rate): ", futureRealValue)

	// fmt.Printf("Future Value: %.1f\n", futureValue)
	// fmt.Printf("Future Value (Adjust it for inflation rate): %.1f\n", futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjust it for inflation rate): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}
