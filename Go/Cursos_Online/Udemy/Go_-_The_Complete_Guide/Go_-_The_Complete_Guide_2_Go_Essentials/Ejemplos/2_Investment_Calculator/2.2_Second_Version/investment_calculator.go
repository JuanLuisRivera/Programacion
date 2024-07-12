package main

import (
	"fmt"
	"math"
)

func main() {
	var investAmout, years float64 = 1000, 10 //We declare both variables on the same line and specify that both variables will be "float64"
	//As well as we use the comma to separate the values each variable will be initialized with, "investAmount" will be "1000" and "years" will be "10"

	expectedReturn := 5.5

	futureValue := investAmout * math.Pow(1+expectedReturn/100, years)

	fmt.Println(futureValue)
}
