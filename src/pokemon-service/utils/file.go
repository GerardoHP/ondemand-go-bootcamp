package utils

import (
	"bufio"
	"context"
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

func (f file) ReadAllFileConcurrent(even bool, items, items_per_worker int) ([]string, error) {
	const workers = 5
	wg := sync.WaitGroup{}
	linesCh := make(chan string)
	var processedLines []string
	mt := sync.Mutex{}
	context, cancel := context.WithCancel(context.Background())
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go processLines(linesCh, &processedLines, &wg, &mt, even, items, items_per_worker, i, cancel)
	}

	wg.Add(1)
	go processFile(f.filename, &wg, linesCh, context)

	wg.Wait()
	cancel()
	return processedLines, nil
}

func processFile(fileName string, wg *sync.WaitGroup, lines chan<- string, ctx context.Context) {
	defer wg.Done()
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ctx.Err() != nil {
			break
		}

		lines <- scanner.Text()
	}
	// if err = scanner.Err(); err != nil {
	// 	return nil, err
	// }
	close(lines)
}

func processLines(lines <-chan string, processedLines *[]string, wg *sync.WaitGroup, mt *sync.Mutex, even bool, items, items_per_worker, worker int, cancel context.CancelFunc) {
	defer wg.Done()
	counter := 0
	for line := range lines {
		fmt.Println("Processing line, ", line, ", from worker ", worker)
		args := strings.Split(line, ",")
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Print("error converting id: ", args[0])
		}

		isEven := id%2 == 0
		if isEven == even {
			mt.Lock()
			currentCount := len(*processedLines)
			mt.Unlock()
			if currentCount >= items {
				fmt.Println("Cancelling work from work ", worker, ", worked ", counter, " items, reached ", currentCount)
				cancel()
				break
			}

			if counter >= items_per_worker {
				fmt.Println("Work completed for worker ", worker, ", worked ", counter, " items, reached ", currentCount)
				break
			}

			mt.Lock()
			*processedLines = append(*processedLines, line)
			mt.Unlock()
			counter++
		}
	}
}
