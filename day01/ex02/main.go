package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

//func main() {
//	oldFS := flag.String("old", "", "old database")
//	newFS := flag.String("new", "", "new database")
//	flag.Parse()
//
//	if *oldFS == "" || *newFS == "" {
//		fmt.Println("usage: ./compareFS --old snapshot1.txt --new snapshot2.txt")
//		return
//	}
//
//	new, err := os.ReadFile(string(*newFS))
//	if err != nil {
//		fmt.Println("cannot read input file:", *newFS, err)
//		return
//	}
//	for _, line := range strings.Split(strings.TrimSuffix(string(new), "\n"), "\n") {
//		fmt.Println(" -> ", line)
//	}
//}

var wg sync.WaitGroup

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
	file1, err := os.ReadFile(string(path1))
	for _, line := range strings.Split(strings.TrimSuffix(string(file1), "\n"), "\n") {
		//fmt.Println(" -> ", line)
		wg.Add(1)
		go func() {
			for s.Scan() {
				//fmt.Println("\t", line, " ? ", s.Text())
				if line == s.Text() {
					wg.Done()
					return
				}
			}
			fmt.Println(message, line)
			wg.Done()
		}()
	}
	//for s.Scan() {
	//	wg.Add(1)
	//	go func() {
	//		for _, line := range strings.Split(strings.TrimSuffix(path1, "\n"), "\n") {
	//			time.Sleep(100)
	//			fmt.Println("\t", line, " ? ", s.Text())
	//			if line == s.Text() && s.Text() != "" {
	//				wg.Done()
	//				return
	//			}
	//		}
	//		if s.Text() != "" {
	//			fmt.Println(message, s.Text())
	//		}
	//		wg.Done()
	//	}()
	//	fmt.Println(message, " -> ", s.Text())
	//}
	wg.Wait()
	err = s.Err()
	if err != nil {
		log.Fatal(err)
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

	compare(*oldFS, *newFS, "ADDED")
	//compare(*newFS, *oldFS, "REMOVED")
}
