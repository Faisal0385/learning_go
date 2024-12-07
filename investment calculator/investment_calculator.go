package main

import (
	"fmt"
	"math"
)

func mainx() {
	const inflationRate = 6.5

	// getting investemet amount value from user
	var investemetAmount float64
	fmt.Print("Investemet Amount: ")
	fmt.Scan(&investemetAmount)

	// getting expected return rate value from user
	var expectedReturnRate float64
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// getting years value from user
	var years float64
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investemetAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println(investemetAmount, expectedReturnRate, years)
	fmt.Println(futureValue, futureRealValue)
}
