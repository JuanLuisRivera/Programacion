package main

func main() {

}

func add(a, b interface{}) { //We don not want to specify the use of the function "add" to a certain data type
	//So we use an interface to be able to make it more generic
	return a + b //The problem is that as the data is too generic, there is no possible way to ensure all data types will
	//Have a valid "+" operator, for example a "struct" + "struct" does not have a defined result
}
