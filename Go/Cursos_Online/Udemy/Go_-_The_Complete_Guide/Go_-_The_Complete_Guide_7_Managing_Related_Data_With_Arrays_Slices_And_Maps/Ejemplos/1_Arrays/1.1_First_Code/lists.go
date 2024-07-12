package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string          // We declare the variable "productNames" as an empty set that must contain 5 elemnts
	productNames = [4]string{"Prueba1"} // We set the value of the first element in the array

	prices := [4]float64{1.01, 19.2, 13.1, 0.5} // We declare a variable prices as a "list" of "4" values of type "float64" with the values
	//"1.01", "19.2", "13.1", "0.5" and we store the data on the variable "prices"

	newPrices := []float64{1.5}
	discountPrices := []float64{100.03, 54.2, 23.1}  // Definimos un nuevo "slice"
	newPrices = append(newPrices, discountPrices...) // Ocupamos el operador "..." para extraer los elementos del slice "discountPrices",
	// De manera que podamos utilizar la funcion "append" para agregar los valores sin generar un error.
	fmt.Println(newPrices)

	fmt.Println(prices)       // We print the data with the stored values
	fmt.Println(productNames) // We print the array

	fmt.Println(prices[2]) // We print the third element of the "prices" array

	productNames[2] = "Prueba3" // We set the single value of the third element in the array
	fmt.Println(productNames)

	featuredPrices := prices[1:3] // We define a "slice" as a subset of values of the array "featuredPrices" with the operator ":"
	fmt.Println(featuredPrices)

	featuredPrices2 := prices[:3] // We define a slice without specifying the lower boundary, so Go creates the slice starting from the first
	// Element in the original array
	fmt.Println(featuredPrices2)

	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices) // We create a slice using another slice as the original data source

	featuredPrices2[0] = 0.0 // We modify the first element of the slice, and as slices are "linked" to arrays, the same data is modified on all the arrays
	// As well as the slices the element is part of
	fmt.Println(featuredPrices)
	fmt.Println(prices)

	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // We output the lenght of the slice and its capacity
	// The lenght of the array is equivalent to the number of elements the array has, the capacity is a way of referencing the number of posible data the
	// Slice may have access to, and so it can be used to resize the slice from itself
	fmt.Println(highlightedPrices)

	highlightedPrices = highlightedPrices[:3] // We resize the slice from itself so the array must have now 3 elements
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	fmt.Println(highlightedPrices) // We compare the resized slice and we can see the complete data we used to have is there, while the earlier slice did not
	// Its important to note that this "resizing" only works with elements to the "right" and never with elements to the left

	dynamicArray := []float64{6.23, 10.9} // If we do not specify the number of elements in the array, GO will automatically create an array, and from there it will
	// Create the slice for the dynamic array, that will be deleted the moment the data size increases
	fmt.Println(dynamicArray)

	newdynamicArray := append(dynamicArray, 5.99) // We add a new element to the slice using the "append" function
	fmt.Println(newdynamicArray)
}
