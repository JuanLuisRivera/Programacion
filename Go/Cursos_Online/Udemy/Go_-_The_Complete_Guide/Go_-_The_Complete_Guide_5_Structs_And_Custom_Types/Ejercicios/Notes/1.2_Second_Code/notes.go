package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "example.com/notes/note"
)

func main() {
	userNote, err := note.New(getNoteData())

	if err != nil {
		fmt.Println(err)
		return
	}

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

	reader := bufio.NewReader(os.Stdin) //Scan does not work well for more than 1 character input so instead it will be used
	//The "bufio" method of the "bufio" package

	text, err := reader.ReadString('\n') //We identify the last character that will be taken in consideration for the input

	if err != nil {
		return "", err
	}

	text = strings.TrimSuffix(text, "\n") //We remove the line break character from the input using the
	//"TrimSuffix" function of the "strings" package

	text = strings.TrimPrefix(text, "\r") //In windows the line break is created with the '\r\n', so it necessary to trim both characters

	return text, nil
}

func getNoteData() (string, string) {
	title, _ := getUserInput("Note title:")

	content, _ := getUserInput("Note content:")

	return title, content
}
