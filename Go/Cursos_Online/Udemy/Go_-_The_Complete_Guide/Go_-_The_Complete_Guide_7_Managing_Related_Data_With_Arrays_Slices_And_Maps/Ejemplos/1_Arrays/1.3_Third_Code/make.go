package main

import "fmt"

type floatMap map[string]float64 // We define a Type Alias to shorten the explicit type

func (m floatMap) output() { // We create a custom method to print the floatMap data
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5) // The "make" function creates a slice with a capacity of "5" and initializes "2" elements of it

	userNames[0] = "Max1"
	userNames[1] = "Max2"

	userNames = append(userNames, "Max3") //We add more data to the slice
	userNames = append(userNames, "Max4")
	userNames = append(userNames, "Max5")
	fmt.Println(&userNames[0])
	fmt.Println(&userNames[4])

	userNames = append(userNames, "Max6") //We can add more data to the slice even if we specified the capacity to be of 5, but the array underlying
	// will be a different one
	fmt.Println(&userNames[0])
	fmt.Println(&userNames[4])

	// Note: The make function allows us to create "semi fixed" arrays, it will allocate memory equal to the ammount of capacity we first declare
	// Making it a bit more efficient in terms of memory usage, it is the "combination" of a slice and an array

	courseRatings := make(floatMap, 3) // We initialize a slice with 3 values for the map, making the memory usage more efficient
	courseRatings["Go"] = 4.0
	courseRatings["Angular"] = 3.4
	courseRatings["React"] = 4.2

	courseRatings.output()

	for index, value := range userNames { // We go trought the whole slice
		fmt.Println("Index: ", index)
		fmt.Println("Values: ", value)
	}

	for key, value := range courseRatings { // We can also use it with maps
		fmt.Println("key: ", key)
		fmt.Println("Values: ", value)
	}
}
