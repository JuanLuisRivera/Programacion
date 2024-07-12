package main

import "fmt"

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	var name customString //We need to declare the "name" variable as a "customString" struct

	name = "Mario"

	name.log() //We call the method of the "customString" struct
}
