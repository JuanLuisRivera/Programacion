package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFIle() (balance float64, nil error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("File not found")
	}
	balanceText := string(data)
	balance, err = strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to convert to float")
	}

	return
}

func main() {
	var accountBalance, err = readBalanceFromFIle()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------------------")
	}

	for {
		displayOptions() //We call the function from the "displayOptions" file that is part of the same package

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
