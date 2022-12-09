package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func crawlWeb(input <-chan string) <-chan *string {
	output := make(chan *string, len(input))
	maxGoroutines := 8
	guard := make(chan struct{}, maxGoroutines)
	for url := range input {
		wg.Add(1)
		url := url
		guard <- struct{}{}
		go func() {
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			res := string(body)
			output <- &res
			//time.Sleep(3 * time.Second)
			//fmt.Println("pack")
			<-guard
			wg.Done()
		}()
	}
	wg.Wait()
	close(output)
	return output
}

func main() {
	size := 20
	input := make(chan string, size)
	for i := 0; i < size; i++ {
		input <- "http://localhost:3333"
	}
	close(input)

	//ctx, cancel := context.WithCancel(context.Background())
	//doneChan := make(chan interface{})
	sign := make(chan os.Signal)
	output := crawlWeb(input)
	signal.Notify(sign, syscall.SIGINT)
	go func() {
		s := <-sign
		_ = s
		fmt.Println("\nMy job here is done!")
		os.Exit(0)
	}()
	for res := range output {
		fmt.Println(len(output), "->\t", *res)
	}
}
