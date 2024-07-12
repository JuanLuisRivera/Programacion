package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "example.com/notes/note"
	"example.com/notes/todo"
)

type saveToFiler interface {
	SaveToFile() error
}

type outputtable interface {
	saveToFiler
	Display()
}

func main() {
	printAnything(1)
	printAnything("Test")

	userNote, err := note.New(getNoteData())

	todoText, err := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println("Saving to file failed")
		return
	}

	fmt.Println("Todo succesfully saved")

	err = outputData(userNote)

	if err != nil {
		return
	}

	printAnything(todo)

}

func printAnything(value interface{}) {
	fmt.Println(value)

	intVal, ok := value.(int) //We analize if the "value" is of type "int", if true, the variable "ok" will be true, else it will be false and we assign
	//The data to "intValue"

	if ok { //We ensure if the data is of type "int"
		intVal = intVal + 1 //We incrase the value of the "intValue" variable by 1
		fmt.Println("Typed variable: ", intVal)
		return
	}

	floatVal, ok := value.(float64) //We declare the "floatVal" variable to be of type "float" and "ok" to be of type boolean

	if ok { //We ensure if the data is of type "float"
		floatVal = floatVal + 0.5 //We incrase the value of the "floatVal" variable by 0.5
		fmt.Println("Typed variable: ", floatVal)
		return
	}
}

func saveData(data saveToFiler) error {
	err := data.SaveToFile()

	if err != nil {
		fmt.Println("Saving to file failed")
		return err
	}

	fmt.Println("Note succesfully saved")

	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func getUserInput(promtp string) (string, error) {
	fmt.Print(promtp)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	text = strings.TrimSuffix(text, "\n")

	text = strings.TrimPrefix(text, "\r")

	return text, nil
}

func getNoteData() (string, string) {
	title, _ := getUserInput("Note title:")

	content, _ := getUserInput("Note content:")

	return title, content
}
