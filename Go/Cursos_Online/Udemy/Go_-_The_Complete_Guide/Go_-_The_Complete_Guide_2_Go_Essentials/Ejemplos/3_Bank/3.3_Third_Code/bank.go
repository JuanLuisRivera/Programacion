package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome")
	fmt.Println("What do you want to do?")
	fmt.Println("1.- Check Balance")
	fmt.Println("2.- Deposit Money")
	fmt.Println("3.- Withdraw Money")
	fmt.Println("4.- Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	wantsCheckBalance := choice == 1

	if wantsCheckBalance {
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 {
		var depositAmount float64

		fmt.Print("Your deposit: ")
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount")
			return //It avoids the execution of any on the following statements
		}

		accountBalance = accountBalance + depositAmount

		fmt.Println("Balance: ", accountBalance)
	} else if choice == 3 {
		var withdrawAmount float64

		fmt.Print("Amount to withdraw<: ")
		fmt.Scan(&withdrawAmount)

		if withdrawAmount > accountBalance {
			fmt.Println("There is not enough money to withdraw")
			return
		}

		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount")
			return
		}

		accountBalance = accountBalance - withdrawAmount

		fmt.Println("Balance: ", accountBalance)
	} else {
		fmt.Print("Bye")
	}
}
