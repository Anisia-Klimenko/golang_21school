package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func multiplex(in ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	for _, c := range in {
		wg.Add(1)
		c := c
		go func() {
			for d := range c {
				out <- d
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	//var i any
	//var count int = 5
	//i = count
	//fmt.Println(i)
	//i = "Hello World!!"
	//fmt.Println(i)

	ch1 := make(chan any, 3)
	ch1 <- 1
	ch1 <- "2"
	ch1 <- 3.0
	close(ch1)

	ch2 := make(chan interface{}, 3)
	ch2 <- 1
	ch2 <- "4"
	ch2 <- 3.1
	close(ch2)
	res := multiplex(ch1, ch2)
	for {
		d, ok := <-res
		if !ok {
			break
		}
		fmt.Println(d)
	}

}
