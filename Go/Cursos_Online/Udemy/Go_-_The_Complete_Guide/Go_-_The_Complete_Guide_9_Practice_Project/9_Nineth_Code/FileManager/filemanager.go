package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (paths FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(paths.InputFilePath)
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

func (paths FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(paths.OutputFilePath)

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

func New(input string, outup string) *FileManager {
	return &FileManager{
		InputFilePath:  input,
		OutputFilePath: outup,
	}
}
