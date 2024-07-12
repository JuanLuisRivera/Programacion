package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "example.com/notes/note"
	"example.com/notes/todo"
)

type saveToFiler interface { //We define an interface in order to avoid the function duplication of the structs "note" and "todo"
	SaveToFile() error //We assure that the struct that uses the interface must have a method, in this case "Save" and that method will
	//Take no arguments an will return an error
}

func main() {
	userNote, err := note.New(getNoteData())

	todoText, err := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil { //We will terminate the program in case an error occurs
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo) //We use the interface with our "saveData" function

	if err != nil {
		fmt.Println("Saving to file failed")
		return
	}

	fmt.Println("Todo succesfully saved")

	userNote.Display()

	err = saveData(userNote)

	if err != nil {
		return
	}

}

func saveData(data saveToFiler) error { //We define a function to save the data of the structs using the interface we created
	err := data.SaveToFile()

	if err != nil {
		fmt.Println("Saving to file failed")
		return err
	}

	fmt.Println("Note succesfully saved")

	return nil
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
