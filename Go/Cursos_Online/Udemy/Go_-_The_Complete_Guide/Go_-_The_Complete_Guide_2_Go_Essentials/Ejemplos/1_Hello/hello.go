package main //Main package, the "main" name is a convention that lets the system know that this package
			 //Should be compiled as an executable instead of a library

import "fmt" //Import the format library that takes care of the I/O

func main() { //Main function, the "main" name is a convention that lets the system know that this is the main function
			  //and is the starting point of the program
	fmt.Print("Hello World!") //Function, part of the Go standard library "fmt", that prints the text "Hello World!"
}
