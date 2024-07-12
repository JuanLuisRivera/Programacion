package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) { //We create the constructor for the "Note" struct
	if title == "" || content == "" { //We return an empty string in case no input is provided
		return Note{}, errors.New("no input received")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil //We return a new "Note" and a "nil" error
}

func (note Note) Display() {
	fmt.Printf("Your note %v, has the following content: \n\n %v", note.title, note.content)
}
