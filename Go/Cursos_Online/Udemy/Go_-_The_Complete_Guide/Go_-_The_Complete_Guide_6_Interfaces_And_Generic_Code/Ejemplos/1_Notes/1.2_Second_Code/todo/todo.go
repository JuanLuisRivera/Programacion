package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(Content string) (Todo, error) {
	if Content == "" {
		return Todo{}, errors.New("invalid input for todo")
	}

	return Todo{
		Text: Content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func (todo Todo) SaveToFile() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
