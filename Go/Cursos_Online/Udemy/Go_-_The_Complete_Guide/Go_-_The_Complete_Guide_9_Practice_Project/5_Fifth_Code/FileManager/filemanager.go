package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("failed to read line in file")
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	file.Close()
	return nil
}
