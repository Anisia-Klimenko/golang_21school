package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if idx >= len(arr) {
		var ErrOutOfRange = errors.New("index " + string(rune(idx)) + " is out of range")
		return 0, ErrOutOfRange
	}
	if idx < 0 {
		var ErrNegativeIndex = errors.New("index cannot be negative")
		return 0, ErrNegativeIndex
	}
	start := unsafe.Pointer(&arr[0])
	size := unsafe.Sizeof(int(0))
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
	}
	fmt.Printf("arr[%d] = %d\n\n", 0, res)

	res, err = getElement(vals, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("arr[%d] = %d\n\n", 4, res)

	res, err = getElement(vals, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("arr[%d] = %d\n\n", 3, res)

	res, err = getElement(vals, -1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("arr[%d] = %d\n\n", -1, res)

}
