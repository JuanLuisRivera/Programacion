package note

import (
	"encoding/json" //Package that will allow to generate a json file with the data gathered
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("no input received")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note %v, has the following content: \n\n %v \n\n", note.title, note.content)
}

func (note Note) SaveToFile() error {
	fileName := strings.ReplaceAll(note.title, " ", "_") //We replace all blank spaces with underscore characters
	fileName = strings.ToLower(fileName) + ".json"       //We convert all uppercase characters to lower case characters and concatenate
	//Json file extension to the filename

	json, err := json.Marshal(note) //We create a file using the "Marshal" method of the json package

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
