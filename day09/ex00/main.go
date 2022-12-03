package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sleepSort(array []int) chan int {
	c := make(chan int)
	for _, elem := range array {
		elem := elem
		fmt.Println("elem:", elem)
		//wg.Add(1)
		go func() {
			time.Sleep(time.Duration(elem * 10))
			c <- elem
			//wg.Done()
		}()
		//return c
	}
	return c
}

func printInt(c chan int) {
	fmt.Println("var:", <-c)
}

func main() {
	array := []int{5, 3, 4, 2, 3, 1}
	c := make(chan int)
	go func() {
		c = sleepSort(array)
	}()
	//for _ = range array {
	go func() {
		//wg.Add(1)
		printInt(c)
		//wg.Done()
	}()
	//}

	wg.Wait()
	close(c)
}
