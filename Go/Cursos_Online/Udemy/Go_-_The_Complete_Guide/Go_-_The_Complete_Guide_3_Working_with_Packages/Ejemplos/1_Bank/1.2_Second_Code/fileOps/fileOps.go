package fileOps //We declare that this file will be part of the fileOps package

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, float float64, permision fs.FileMode) { //The function needs to star with capital letter to be propely exported
	//To other packages
	floatToText := fmt.Sprint(float)

	os.WriteFile(fileName, []byte(floatToText), permision)
}

func ReadFloatFromFile(defaultValue float64, fileName string) (float64, error) { //Function name must start with a capital letter to be used on the
	//Main package
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("file not found")
	}

	dataToString := string(data)
	float, err := strconv.ParseFloat(dataToString, 64)

	if err != nil {
		return defaultValue, errors.New("failed to convert to float")
	}

	return float, err
}

//It should be noted that for any variables, name, and functions from external package the same principle applies, in order to use it
//It needs to start with a capital letter else it will not be possible to be used externally
