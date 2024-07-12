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
	// 1) Create a new array (!) that contains three hobbies you have
	// 	Output (print) that array in the command line.
	hobbies := [3]string{"Cocinar\n", "Estudiar\n", "Oir musica\n"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//	- The first element (standalone)
	//	- The second and third element combined as a new list
	fst := hobbies[0]
	fmt.Println(fst)

	hobbies2 := [2]string{hobbies[1], hobbies[2]}
	fmt.Println(hobbies2)

	// 3) Create a slice based on the first element that contains
	//	the first and second elements.
	//	Create that slice in two different ways (i.e. create two slices in the end)
	fstslice := hobbies[0:1]
	sndslice := hobbies[:1]
	fmt.Println(fstslice)
	fmt.Println(sndslice)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//	and last element of the original array.
	thdslice := append(sndslice, hobbies[2]) // Re - size of the slice
	fmt.Println(thdslice)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Aprender a programar en GO", "Integrar GO a desarrollo web"}
	fmt.Println(goals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Mejorar mi conocimiento de lenguajes de pregramacion"
	goals = append(goals, "Adquirir mas experiencia para mejorar mi habilidad de programacion")
	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//	dynamic list of products (at least 2 products).
	//	Then add a third product to the existing list of products.
	prod1 := Product{
		title: "Manzana",
		id:    "0",
		price: 1.05,
	}

	prod2 := Product{
		title: "Pera",
		id:    "1",
		price: 2.50,
	}

	prod3 := Product{
		title: "Mango",
		id:    "2",
		price: 5.15,
	}

	products := []Product{prod1, prod2}
	fmt.Println(products)

	products = append(products, prod3)
	fmt.Println(products)
}
