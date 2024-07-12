package main

import ( //This import structure is the most common to import Golang libraries
	"fmt"
	"math" //Library that contains the Pow function
)

func main() {
	var investAmout = 1000   //Variable definition and initialization
	var expectedReturn = 5.5 //Variable definition and initialization
	var years = 10           //Variable definition and initialization

	var futureValue = float64(investAmout) * math.Pow(1+expectedReturn/100, float64(years)) //As Go is an static typed language we need to
	//Convert the Int type variable on to float64
	//To be able to perform the calculation

	fmt.Println(futureValue) //We print the result of the calculation to the console
}
