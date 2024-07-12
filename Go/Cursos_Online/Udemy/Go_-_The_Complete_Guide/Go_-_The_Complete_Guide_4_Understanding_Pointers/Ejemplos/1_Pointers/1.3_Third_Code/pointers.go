package main

import "fmt"

func main() {
	age := 32
	agePointer := &age

	fmt.Println("Age address in memory: ", agePointer)
	fmt.Println("Age: ", *agePointer)

	getAdultYears(*&agePointer)
	fmt.Println("Adult years: ", age) //As we modified the value of the variable, we expect it to become 14 instead of 32
}

func getAdultYears(agePointer *int) {
	*agePointer = *agePointer - 18 //We modify the value of the variable instead of modifying a copy if the variable
}
