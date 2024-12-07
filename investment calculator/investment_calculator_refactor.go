package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

	// getting investemet amount value from user
	investemetAmount := outputText("Investemet Amount: ")

	// getting expected return rate value from user
	expectedReturnRate := outputText("Expected Return Rate: ")

	// getting years value from user
	years := outputText("Years: ")

	futureValue, futureRealValue := calculteFutureValue(investemetAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Inflation): %.2f", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)

}

func outputText(label string) float64 {
	var userInput float64
	fmt.Print(label)
	fmt.Scan(&userInput)
	return userInput
}

func calculteFutureValue(investemetAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investemetAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow((1+inflationRate/100), years)
	return fv, rfv
}
