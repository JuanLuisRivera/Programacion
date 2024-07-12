package main

import "fmt"

func main() {
	result := add(1, 2) //We no longer have problems with out function as it is capable to infer the correct data type
	fmt.Println(result)
}

func add[T int | float64 | string](a, b T) T { //We convert the function to a "generic" function by the use of "[]" after declare
	//The name, and we specify the letter "T" as a placeholder of either of the data type "int", "float64" and "string"

	//By using this generic function we don not need to use data validation as "Go" automatically infers the type of data
	//The variable has

	return a + b

}
