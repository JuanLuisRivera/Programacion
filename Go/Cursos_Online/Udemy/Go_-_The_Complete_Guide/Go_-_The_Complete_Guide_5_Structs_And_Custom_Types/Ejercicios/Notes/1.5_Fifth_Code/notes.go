package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "example.com/notes/note"
	"example.com/notes/todo"
)

func main() {
	userNote, _ := note.New(getNoteData())

	todoText, _ := getUserInput("Todo text: ")

	todo, err := todo.New(todoText) //We create "todo" struct

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.SaveToFile()

	if err != nil {
		fmt.Println("Saving to file failed")
		return
	}

	fmt.Println("Todo succesfully saved")

	userNote.Display()

	err = userNote.SaveToFile()

	if err != nil {
		fmt.Println("Saving to file failed")
		return
	}

	fmt.Println("Note succesfully saved")
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
