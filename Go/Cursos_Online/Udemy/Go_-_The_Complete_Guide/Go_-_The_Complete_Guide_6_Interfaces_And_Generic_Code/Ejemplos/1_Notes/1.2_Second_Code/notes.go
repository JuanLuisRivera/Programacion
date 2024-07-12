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

type outputtable interface { //We create an embedded interface called "outputtable" that increases the methods defined by the "saveToFiler" interface
	//So that more data can actually be used with the interfaces and therefore diminish the ammount of replication on the main function
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

	printAnything(todo) //We can pass the "todo" data type but it should display the default case

}

func printAnything(value interface{}) { //We use the "interface {}" declaration to specify that all types are allowed, alternatively the keyword "any" can be used
	//As well instead of the "interface {} declaration"
	fmt.Println(value)

	switch value.(type) { //Podemos generar tambien un switch que analice los valores de los datos recibidos y ejecute una funcion especifica de acuerdo al tipo
	//De dato que le pasamos a la funcion
	case int: //En caso de que el valor sea un entero
		fmt.Println("Integer: ", value)
	case float64: //Caso donde el tipo de dato es flotante
		fmt.Println("Float: ", value)
	case string: //Caso donde se recibe una cadena
		fmt.Println(value)
	default: //Añadimos un caso por default
		fmt.Println("Algun tipo de dato no reconocido")
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
