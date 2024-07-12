package main

import (
	"fmt"
	"os"
	"strconv" //Package that allows to use convertions from string to other data types
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFIle() (balance float64) { //Function that reads a file and returns the value as a float
	data, _ := os.ReadFile(accountBalanceFile) //As the function returns a byte array and an error, we have to assing both values to a
	//Variable each, so we use the "_" as a place to ignore the variable that we will not use
	balanceText := string(data)                      //We use the array of bytes and convert them to a string
	balance, _ = strconv.ParseFloat(balanceText, 64) //We use the "ParseFloat" function to convert the string we read from the file to float
	//And the use the "_" to ingore the error value the function also returns
	return
}

func main() {
	var accountBalance float64 = readBalanceFromFIle() //It should be noted that the file should already exist, else an error will occur

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
