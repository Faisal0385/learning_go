package main

import "fmt"

func mainx() {
	var revenue float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	var expenses float64
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	var taxRate float64
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	profit := EBT * (1 - taxRate/100)

	ratio := EBT / profit

	fmt.Println(EBT)
	fmt.Println(profit)
	fmt.Println(ratio)
}
