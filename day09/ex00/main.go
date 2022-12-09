package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sleepSort(array []int) <-chan int {
	c := make(chan int, len(array))
	for _, elem := range array {
		wg.Add(1)
		elem := elem
		go func() {
			time.Sleep(time.Duration(elem))
			c <- elem
			wg.Done()
		}()
	}
	wg.Wait()
	close(c)
	return c
}

func main() {
	array := []int{5, 3, 4, 2, 3, 1}
	c := sleepSort(array)
	for d := range c {
		fmt.Println(d)
	}
}
