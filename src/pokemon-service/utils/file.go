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
	const workers = 1
	wg := sync.WaitGroup{}
	linesCh := make(chan string)
	pokemonCount := make(chan int)
	mt := sync.RWMutex{}
	file, err := os.Open(f.filename)
	if err != nil {
		log.Panic("Failed to open", err)
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go readLine(scanner, even, items, items_per_worker, linesCh, pokemonCount, &wg, &mt)
	}

	wg.Add(1)
	go coordinateChan(&lines, &wg, linesCh, pokemonCount)
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	wg.Wait()
	return lines, nil
}

func coordinateChan(lines *[]string, wg *sync.WaitGroup, linesCh <-chan string, pokemonCount chan<- int) {
	defer wg.Done()
	for line := range linesCh {
		*lines = append(*lines, line)
		pokemonCount <- len(*lines)
	}

	close(pokemonCount)
}

func readLine(s *bufio.Scanner, even bool, items, items_per_worker int, linesCh chan<- string, pokemonCount <-chan int, wg *sync.WaitGroup, mt *sync.RWMutex) {
	defer wg.Done()
	counter := 0
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
		nPok := <-pokemonCount
		if isEven == even {
			linesCh <- line
			counter++
			if counter == items_per_worker || nPok > items {
				break
			}
		}
	}

	close(linesCh)
}
