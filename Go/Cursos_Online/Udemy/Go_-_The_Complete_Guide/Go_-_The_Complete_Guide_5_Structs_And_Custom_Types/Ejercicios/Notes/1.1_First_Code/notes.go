package main

import (
	"fmt"

	note "example.com/notes/note"
)

func main() {
	userNote, err := note.New(getUserInput("Note title:"), getUserInput("Note content:")) //We create the "userNote" by passing the appropiate
	//Arguments

	if err != nil {
		println(err)
		return
	}

	userNote.Display()

}

func getUserInput(promtp string) string {
	fmt.Print(promtp)

	var data string
	fmt.Scanln(&data) //We use "Scanln" as a way to ensure the input only is taken after te key "enter" has been pressed

	return data
}

func getNoteData() (string, string) { //Function that obtain the data for the note app
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}
