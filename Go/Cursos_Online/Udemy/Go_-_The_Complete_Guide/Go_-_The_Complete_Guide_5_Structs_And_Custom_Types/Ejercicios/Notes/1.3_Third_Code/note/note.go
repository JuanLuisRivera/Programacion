package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string //We capitalize the variables in order to use the "Marshall" function
	Content   string
	CreatedAt time.Time
}

func New(Title, Content string) (Note, error) {
	if Title == "" || Content == "" {
		return Note{}, errors.New("no input received")
	}

	return Note{
		Title:     Title,
		Content:   Content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note %v, has the following content: \n\n %v \n\n", note.Title, note.Content)
}

func (note Note) SaveToFile() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note) //The "Marshal" function can only convert public data so we have to capitalize the variables inside the struct

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
