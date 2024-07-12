package main

func main() {
	result := add(1, 2)

	result = result + 1 //We receive an error as the return data is too generic to be used and exist a mismatch
	//Between data types on of type "any" and the other one of type "int"
}

func add(a, b interface{}) any { //We generalize the data return by declaring it will return "any" datatype
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt { //We can ensure the code works for the "int" datatype but it will be too restrictive
		//When returning the value
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat { //We have to ensure the data return can accept a "float" and an "int" datatype
		return aFloat + bFloat
	}

	aString, aIsString := a.(float64)
	bString, bIsString := b.(float64)

	if aIsString && bIsString { //We face no problem as we have already specified the data return type will be "any"
		return aString + bString
	}

	return nil

}
