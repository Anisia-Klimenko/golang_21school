package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

//func crawlWeb(input <-chan string) <-chan *string {
//	output := make(chan *string, len(input))
//	maxGoroutines := 8
//	guard := make(chan struct{}, maxGoroutines)
//	for url := range input {
//		wg.Add(1)
//		url := url
//		guard <- struct{}{}
//		go func() {
//			resp, err := http.Get(url)
//			if err != nil {
//				log.Fatalln(err)
//			}
//			body, err := ioutil.ReadAll(resp.Body)
//			if err != nil {
//				log.Fatalln(err)
//			}
//			res := string(body)
//			output <- &res
//			//time.Sleep(3 * time.Second)
//			//fmt.Println("pack")
//			<-guard
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//	close(output)
//	return output
//}

var mutex = sync.Mutex{}
var cors = 0

func crawlWeb(inCh <-chan string, ctx context.Context) <-chan *string {
	guard := make(chan struct{}, 8)

	outCh := make(chan *string, len(inCh))
	defer close(outCh)
	for url := range inCh {
		time.Sleep(time.Duration(1 * int(time.Second)))
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				mutex.Lock()
				cors++
				fmt.Printf("%d goroutine started\n", cors)
				mutex.Unlock()

				time.Sleep(time.Duration(3 * int(time.Second)))

				guard <- struct{}{}
				resp, err := http.Get(url)
				if err != nil {
					log.Fatalln(err)
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
				res := string(body)
				outCh <- &res

				<-guard
				fmt.Println("goroutine ended")
			}
		}(url)
	}
	wg.Wait()
	return outCh
}

func main() {
	size := 20
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	input := make(chan string, size)
	for i := 0; i < size; i++ {
		input <- "http://localhost:3333"
	}
	close(input)

	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("\nMy job here is done!")
		cancel()
	}()
	outCh := crawlWeb(input, ctx)
	for res := range outCh {
		fmt.Println(len(outCh), "->\t", *res)
	}
}
