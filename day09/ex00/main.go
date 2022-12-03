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
		//wg.Add(1)
		go func() {
			time.Sleep(time.Duration(elem * 10))
			fmt.Println("elem:", elem)
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
	c := make(chan int, len(array))
	go func() {
		c = sleepSort(array)
	}()
	//for _ = range array {
	//go func() {
	//	//wg.Add(1)
	//	printInt(c)
	//	//wg.Done()
	//}()
	//}

	for i := range c {
		println(i)
	}

	wg.Wait()
	close(c)
}
