package main

import "fmt"

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.2f\n", ratio)
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (EBT float64, profit float64, ratio float64) {
	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = EBT / profit

	return EBT, profit, ratio
}
