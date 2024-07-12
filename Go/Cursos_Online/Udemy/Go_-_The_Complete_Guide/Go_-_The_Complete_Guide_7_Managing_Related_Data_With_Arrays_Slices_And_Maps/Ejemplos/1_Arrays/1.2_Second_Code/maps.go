package main

import "fmt"

type Product struct { // Structs are used whenever the data is clearly defined
	id    string
	title string
	price float64
}

// Note: Structs are better suited for data entities such as users, cars, buildings, items, while maps are better for a collection of values that is not well defined

func main() {
	websites := []string{"https://google.com", "https://aws.com"} // We create an array of websites urls

	websitesMap := map[string]string{"Google": "https://google.com", "Amazon Web Services": "htpps://aws.com"} //We create a map with the keys "[string]" and values "{string}"

	fmt.Println(websites)
	fmt.Println(websitesMap)

	fmt.Println(websitesMap["Amazon Web Services"]) // We print specific data with only the key value

	websitesMap["LinkedIn"] = "https://linkedin.com" // We append new data to the map with only a key that does not exist and an associated value
	fmt.Println(websitesMap)

	websitesMap["Google"] = "https://8.8.8.8" // We override the value associated with the specific key
	fmt.Println(websitesMap)

	delete(websitesMap, "Google") // We delete the key with the "delete" function
	fmt.Println(websitesMap)
}
