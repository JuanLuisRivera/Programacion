package main

import (
	"fmt"
	"os" //We import the OS pakage that allows to intereact with the filesystem
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance) //We convert the balance variable to a string using the "Sprint" function

	os.WriteFile("balance.txt", []byte(balanceText), 0644) //We create the file, transform the converted String to a "collection"
	//Of bytes and set the permission by the User, the Group and Everyone else
}

func main() {
	var accountBalance float64 = 1000

	for {
		fmt.Println("Welcome")
		fmt.Println("What do you want to do?")
		fmt.Println("1.- Check Balance")
		fmt.Println("2.- Deposit Money")
		fmt.Println("3.- Withdraw Money")
		fmt.Println("4.- Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
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

			writeBalanceToFile(accountBalance)

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

			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Bye")
			fmt.Println("Thanks for coming by")
			return
		}

	}

}
