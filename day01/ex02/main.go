package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func compare(path1 string, path2 string, message string) {
	file2, err := os.Open(path2)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file2.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(file2)
	file1, err := os.ReadFile(path1)
	result := make(map[string]bool)
	var found bool

	for s.Scan() {
		for _, line := range strings.Split(strings.TrimSuffix(string(file1), "\n"), "\n") {
			line := line
			found = false
			if line == s.Text() {
				found = true
				break
			}
		}
		if s.Text() != "" && !found {
			result[s.Text()] = true
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	printResult(message, result)
}

func printResult(message string, result map[string]bool) {
	for s := range result {
		fmt.Println(message, s)
	}
}

func main() {
	oldFS := flag.String("old", "", "old database")
	newFS := flag.String("new", "", "new database")
	flag.Parse()

	if *oldFS == "" || *newFS == "" {
		fmt.Println("usage: ./compareFS --old snapshot1.txt --new snapshot2.txt")
		return
	}

	compare(*newFS, *oldFS, "ADDED")
	compare(*oldFS, *newFS, "REMOVED")
}
