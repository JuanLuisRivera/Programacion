package main

import (
	"fmt"

	"example.com/bank/fileOps"

	"github.com/Pallinder/go-randomdata" //We export the package from the repository using "go get github.com/Pallinder/go-randomdata" command
	//On the temrinal as well as here
)

const accountBalanceFile = "balance.txt" //We set the filename that will be used
const permision = 0644                   //We set the permision fo the moment the file will be created

func main() {
	var accountBalance, err = fileOps.ReadFloatFromFile(0, accountBalanceFile) //We use the function on the "fileOps" package to read the data
	//And set "accountBalance" to 0 in case the file does not exist

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("Setting account balance to 0")
		fmt.Println("---------------------")
	}

	for {

		fmt.Println("Welcome")
		fmt.Println(randomdata.PhoneNumber()) //We use the function "PhoneNumber" given by the external package "randomdata"

		displayOptions() //We use the function on the "displayOptions" file

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

			fileOps.WriteFloatToFile(accountBalanceFile, accountBalance, permision) //We use the function on the "fileOps" package to write data

		case 3:
			var withdrawAmount float64

			fmt.Print("Amount to withdraw: ")
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

			fileOps.WriteFloatToFile(accountBalanceFile, accountBalance, permision) //We use the function on the "fileOps" package to write data

		default:
			fmt.Println("Bye")
			fmt.Println("Thanks for coming by")
			return
		}

	}

}
