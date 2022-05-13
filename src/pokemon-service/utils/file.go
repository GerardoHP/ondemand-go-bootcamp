package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

type file struct {
	filename string
}

type File interface {
	ReadAllFileLines() ([]string, error)
	AppendLineToFile(string) error
}

// Gets a new instance of the file utils
func NewFileUtil(fn string) File {
	return &file{filename: fn}
}

// Read all the lines in a file
func (f file) ReadAllFileLines() ([]string, error) {
	file, err := os.Open(f.filename)
	if err != nil {
		log.Panic("Failed to open", err)
		return nil, err
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Append a line to the file
func (f file) AppendLineToFile(line string) error {
	lines, err := f.ReadAllFileLines()
	if err != nil {
		return err
	}

	fileContent := ""
	for _, line := range lines {
		fileContent += line
		fileContent += "\n"
	}

	fileContent += line
	err = ioutil.WriteFile(f.filename, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
