package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

var lines = false
var chars = false
var words = false

func main() {
	if !getFlags() {
		fmt.Println("only one flag (l|m|w) is required")
		return
	}
	var wg sync.WaitGroup
	for _, file := range flag.Args() {
		wg.Add(1)
		file := file
		go func() {
			if lines {
				fmt.Printf("%d\t%s\n", countLines(file), file)
			} else if words {
				fmt.Printf("%d\t%s\n", countWords(file), file)
			} else if chars {
				fmt.Printf("%d\t%s\n", countChars(file), file)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func countChars(path string) int {
	file, _ := os.ReadFile(path)
	return len(file)
}

func countWords(path string) int {
	var count int
	fileHandle, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(fileHandle *os.File) {
		err := fileHandle.Close()
		if err != nil {

		}
	}(fileHandle)
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanWords)

	for fileScanner.Scan() {
		count++
	}
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}
	return count
}

func countLines(path string) int {
	var count int
	file, _ := os.ReadFile(path)
	for range strings.Split(strings.TrimSuffix(string(file), "\n"), "\n") {
		count++
	}
	return count
}

func getFlags() bool {
	flag.BoolVar(&lines, "l", false, "count lines")
	flag.BoolVar(&chars, "m", false, "count characters")
	flag.BoolVar(&words, "w", false, "count words")
	flag.Parse()

	fmt.Println("lines", lines)
	fmt.Println("chars", chars)
	fmt.Println("words", words)

	if !lines && !chars && !words {
		words = true
	}
	if (lines && !chars && !words) || (!lines && chars && !words) || (!lines && !chars && words) {
		return true
	}
	return false
}
