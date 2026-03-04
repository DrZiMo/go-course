package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilepath  string
	OutputFilepath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilepath)

	if err != nil {
		return nil, errors.New("Could not open file!")
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Reading file content failed!")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJSON(data any) error {
	file, err := os.Create(fm.OutputFilepath)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert data into json.")
	}

	file.Close()
	return nil
}

func New(inputFilepath, outputFilepath string) *FileManager {
	return &FileManager{
		InputFilepath:  inputFilepath,
		OutputFilepath: outputFilepath,
	}
}
