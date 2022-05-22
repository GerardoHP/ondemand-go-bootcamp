package utils

import (
	"bufio"
	"context"
	"errors"
	"fmt"
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

// Reads a file with a pool of parallel workes, 5 by default
func (f file) ReadAllFileConcurrent(even bool, items, items_per_worker int) ([]string, error) {
	const workers = 5
	wg := sync.WaitGroup{}
	linesCh := make(chan string)
	errorsCh := make(chan string)
	var processedLines []string
	mt := sync.Mutex{}
	context, cancel := context.WithCancel(context.Background())
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go processLines(linesCh, errorsCh, &processedLines, &wg, &mt, even, items, items_per_worker, i, cancel)
	}

	wg.Add(1)
	go processFile(f.filename, &wg, linesCh, errorsCh, context)

	wg.Add(1)
	var err error
	go func(errorCh, lines chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		if stringErr, ok := <-errorCh; ok && stringErr != "" {
			err = errors.New(stringErr)
			close(errorCh)
			close(linesCh)
		}
	}(errorsCh, linesCh, &wg)

	wg.Wait()
	cancel()
	return processedLines, err
}

// Process the file from start to end unless the routine gets cancelled
func processFile(fileName string, wg *sync.WaitGroup, lines, errors chan<- string, ctx context.Context) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		errors <- err.Error()
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ctx.Err() != nil {
			errors <- ""
			break
		}

		lines <- scanner.Text()
	}

	if err = scanner.Err(); err != nil {
		errors <- err.Error()
		return
	}

	close(lines)
}

// Worker that process the lines
func processLines(lines <-chan string, errors chan<- string, processedLines *[]string, wg *sync.WaitGroup, mt *sync.Mutex, even bool, items, items_per_worker, worker int, cancel context.CancelFunc) {
	defer wg.Done()
	counter := 0
	for line := range lines {
		fmt.Println("Processing line, ", line, ", from worker ", worker)
		args := strings.Split(line, ",")
		id, err := strconv.Atoi(args[0])
		if err != nil {
			errors <- "error converting id"
		}

		isEven := id%2 == 0
		if isEven == even {
			mt.Lock()
			*processedLines = append(*processedLines, line)
			currentCount := len(*processedLines)
			mt.Unlock()
			counter++
			if currentCount >= items {
				fmt.Println("Cancelling work from work ", worker, ", worked ", counter, " items, reached ", currentCount)
				cancel()
				break
			}

			if counter >= items_per_worker {
				if currentCount < items {
					continue
				}

				fmt.Println("Work completed for worker ", worker, ", worked ", counter, " items, reached ", currentCount)
				break
			}
		}
	}
}
