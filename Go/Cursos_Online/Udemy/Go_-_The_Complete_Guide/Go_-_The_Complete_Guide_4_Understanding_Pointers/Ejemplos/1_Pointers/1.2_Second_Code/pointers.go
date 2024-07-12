package main

import "fmt"

func main() {
	age := 32
	agePointer := &age //We create a pointer using the "age" address in memory and the "&" operator

	fmt.Println("Age address in memory: ", agePointer) //We print the "age" variable address
	fmt.Println("Age: ", *agePointer)                  //We "dereference" the pointer, making it posible to access the stored value
	//Of the variable "age"

	fmt.Println("Adult years: ", getAdultYears(agePointer)) //This way of creating the programm is useful as it creates
	//The variable "age" only once, and so it will not use as much memory
}

func getAdultYears(agePointer *int) int { //We use the pointer as the argument of the function
	return *agePointer - 18
}
