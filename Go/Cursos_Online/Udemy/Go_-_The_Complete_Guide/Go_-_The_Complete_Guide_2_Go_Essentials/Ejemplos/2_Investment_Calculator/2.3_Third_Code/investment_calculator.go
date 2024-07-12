package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5 //We define a constant value using thes "const" keyword

	var investAmout, years float64 = 1000, 10

	expectedReturn := 5.5

	futureValue := investAmout * math.Pow(1+expectedReturn/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
