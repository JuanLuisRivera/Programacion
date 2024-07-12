package main

import (
	"errors" //Package that allows to manipulate errors for the functions we will define
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
		return 0, errors.New("File not found") //We set the return value to be 1000 in case the file is not found on the system
		//And we create an error text
	}
	balanceText := string(data)
	balance, err = strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to convert to float") //We set the value to 1000 if the data form the file can not be converted and
		//Return the error string that the convertion had failed
	}

	return
}

func main() {
	var accountBalance, err = readBalanceFromFIle()

	if err != nil { //In case of error we display the message and the error that occurred
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------------------")
		panic("Execution can not be completed as the source data could not be found.") //We use the panic funcion to terminate the execution
		//Of the application, it should be noted that its purpose is intended only for very specific and important situations
	}

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
