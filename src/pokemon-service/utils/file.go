package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type file struct {
	filename string
}

type File interface {
	ReadAllFileLines() ([]string, error)
	AppendLineToFile(string) error
	ReadAllFileConcurrent(even bool, items, items_per_worker int) ([]string, error)
}

// Gets a new instance of the file utils
func New(fn string) File {
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

func (f file) ReadAllFileConcurrent(even bool, items, items_per_worker int) ([]string, error) {
	if items_per_worker == 0 {
		items_per_worker = 1
	}

	wg := sync.WaitGroup{}
	ch := make(chan string)
	mt := sync.RWMutex{}
	file, err := os.Open(f.filename)
	if err != nil {
		log.Panic("Failed to open", err)
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	wg.Add(items_per_worker)
	for i := 0; i < items_per_worker; i++ {
		go readLine(scanner, even, ch, &wg, &mt, i)
	}

	for n := range ch {
		lines = append(lines, n)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	wg.Wait()
	return lines, nil
}

func readLine(s *bufio.Scanner, even bool, c chan<- string, wg *sync.WaitGroup, mt *sync.RWMutex, routineId int) {
	print(routineId)
	for s.Scan() {
		mt.RLock()
		line := s.Text()
		mt.RUnlock()
		lineDesc := strings.Split(line, ",")

		id, err := strconv.Atoi(lineDesc[0])
		if err != nil {
			log.Print("failed to convert id")
			continue
		}

		isEven := id%2 == 0
		if isEven == even {
			c <- line
		}
	}

	close(c)
	wg.Done()
}
