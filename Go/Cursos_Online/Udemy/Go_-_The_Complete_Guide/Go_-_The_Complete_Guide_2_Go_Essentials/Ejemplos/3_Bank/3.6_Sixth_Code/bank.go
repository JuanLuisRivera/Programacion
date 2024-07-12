package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	for {
		fmt.Println("Welcome")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		fmt.Println("What do you want to do?")
		fmt.Println("1.- Check Balance")
		fmt.Println("2.- Deposit Money")
		fmt.Println("3.- Withdraw Money")
		fmt.Println("4.- Exit")

		switch choice { //We use "switch" instead of the "if else" statements
		case 1:
			fmt.Println("Your balance is: ", accountBalance)

		case 2:
			var depositAmount float64

			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")

				continue
			}

			accountBalance = accountBalance + depositAmount

			fmt.Println("Balance: ", accountBalance)

		case 3:
			var withdrawAmount float64

			fmt.Print("Amount to withdraw<: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accountBalance {
				fmt.Println("There is not enough money to withdraw")

				continue
			}

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount")

				continue
			}

			accountBalance = accountBalance - withdrawAmount

			fmt.Println("Balance: ", accountBalance)

		default:
			fmt.Print("Bye")
			fmt.Println("Thanks for coming by")
			return //We use "return" again instead of "break" because "break" on a "switch" statement only exits the "switch"
			//Statement instead of exiting the "for" loop, so we need to finish the execution of the application with "return"
		}

	}

}
