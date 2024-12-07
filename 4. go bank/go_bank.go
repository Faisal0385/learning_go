package main

import "fmt"

func mainx() {

	var accountBalance float64 = 1000
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("-------------------")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdrawal money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {

			fmt.Println("Your current balance is: ", accountBalance)

			var DepositAmount float64
			fmt.Print("Your Deposit: ")
			fmt.Scan(&DepositAmount)

			if DepositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			accountBalance += DepositAmount

			fmt.Println("Deposit amount: ", DepositAmount)
			fmt.Println("Updated balance: ", accountBalance)

		} else if choice == 3 {

			fmt.Println("Your current balance is: ", accountBalance)

			var withDrawalValue float64
			fmt.Print("Withdrawal value: ")
			fmt.Scan(&withDrawalValue)

			if withDrawalValue <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withDrawalValue > accountBalance {
				fmt.Println("Invalid amount. You can not withdraw this amount.")
				continue
			}

			accountBalance -= withDrawalValue

			fmt.Println("Withdrawal amount: ", withDrawalValue)
			fmt.Println("Updated balance: ", accountBalance)

		} else if choice == 4 {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Invalid value.")
		}

	}

	fmt.Println("Thank you for choosing us.")
}
