package main

import "fmt"

func main() {
	age := 32 //Regular variable

	fmt.Println("Age: ", age)

	fmt.Println("Adult years: ", getAdultYears(age))
}

func getAdultYears(age int) int { //The variable used in the function has the same value as "age" as it is a copy
	//Doubling the amount of variables in memory when using this approach
	return age - 18
}
