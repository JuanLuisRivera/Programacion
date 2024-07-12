//We create this new file to migrate functions and specify that althought the file is different, is part of the same package

package main //We specify that this file will be part of the main package

import "fmt" //We need to import the "fmt" package in both the "bank" file and here because tha packages
//Does not import althought they might be part of the same package

func displayOptions() { //We define a function that displays all the menu options

	fmt.Println("Welcome")
	fmt.Println("What do you want to do?")
	fmt.Println("1.- Check Balance")
	fmt.Println("2.- Deposit Money")
	fmt.Println("3.- Withdraw Money")
	fmt.Println("4.- Exit")

}
