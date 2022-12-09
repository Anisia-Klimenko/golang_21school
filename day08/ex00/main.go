package main

import (
	"errors"
	"fmt"
	"strconv"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("empty slice")
	} else if idx >= len(arr) {
		return 0, errors.New("index " + strconv.Itoa(idx) + " is out of range")
	} else if idx < 0 {
		return 0, errors.New("index cannot be negative")
	}
	start := unsafe.Pointer(&arr[0])
	size := unsafe.Sizeof(0)
	item := *(*int)(unsafe.Pointer(uintptr(start) + size*uintptr(idx)))
	return item, nil
}

func main() {
	vals := []int{10, 20, 30, 40}
	var res int
	fmt.Println(vals)
	fmt.Println()

	res, err := getElement(vals, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("arr[%d] = %d\n", 0, res)
	}

	res, err = getElement(vals, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("arr[%d] = %d\n", 4, res)
	}

	res, err = getElement(vals, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("arr[%d] = %d\n", 3, res)
	}

	res, err = getElement(vals, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("arr[%d] = %d\n", -1, res)
	}

	res, err = getElement([]int{}, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("arr[%d] = %d\n", -1, res)
	}
}
