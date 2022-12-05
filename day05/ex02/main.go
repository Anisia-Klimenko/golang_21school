package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap struct {
	Presents []Present
}

func (p PresentHeap) Len() int {
	return len(p.Presents)
}

func (p PresentHeap) Less(i, j int) bool {
	if j > p.Len() {
		return true
	} else if i > p.Len() {
		return false
	}
	if p.Presents[i].Value == p.Presents[j].Value {
		return p.Presents[i].Size > p.Presents[j].Size
	}
	return p.Presents[i].Value < p.Presents[j].Value
}

func (p *PresentHeap) Swap(i, j int) {
	if i > p.Len() || j > p.Len() {
		return
	}
	p.Presents[i].Value, p.Presents[j].Value = p.Presents[j].Value, p.Presents[i].Value
	p.Presents[i].Size, p.Presents[j].Size = p.Presents[j].Size, p.Presents[i].Size
}

func (p PresentHeap) isSorted() bool {
	for i := 1; i < p.Len(); i++ {
		if p.Less(i-1, i) {
			return false
		}
	}
	return true
}

func (p *PresentHeap) sort() {
	for i := 0; i < p.Len(); i++ {
		for j := i; j < p.Len(); j++ {
			if p.Less(i, j) {
				p.Swap(i, j)
			}
		}
	}
}

func (p *PresentHeap) Push(x any) {
	p.Presents = append(p.Presents, x.(Present))
	if !p.isSorted() {
		p.sort()
	}
}

func (p *PresentHeap) Pop() any {
	old := p.Presents
	n := len(old)
	if n == 0 {
		return Present{}
	}
	item := old[n-1]
	p.Presents = old[0 : n-1]
	return item
}

func (p PresentHeap) printHeap() {
	for count, pr := range p.Presents {
		fmt.Println(count, " -> value:", pr.Value, "size:", pr.Size)
	}
}

func printSlice(p []Present) {
	for count, pr := range p {
		fmt.Println(count, " -> value:", pr.Value, "size:", pr.Size)
	}
}

func getNCoolestPresents(ps []Present, n int) ([]Present, error) {
	if n < 0 {
		return []Present{}, errors.New("n less than 0")
	}
	if n > len(ps) {
		return []Present{}, errors.New("n too big")
	}
	ph := PresentHeap{}
	for _, pr := range ps {
		ph.Push(pr)
	}
	for ph.Len() > n {
		ph.Pop()
	}
	return ph.Presents, nil
}

func main() {
	parray := []Present{{3, 1}, {4, 5}, {5, 2}}
	ph := PresentHeap{parray}
	heap.Init(&ph)

	fmt.Println("\n====== Heap init ======")
	ph.printHeap()

	fmt.Println("\n==== Get 2 coolest ====")
	res, err := getNCoolestPresents(parray, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		printSlice(res)
	}

	fmt.Println("\n===== Push (5, 1) =====")
	ph.Push(Present{5, 1})
	ph.printHeap()

	fmt.Println("\n==== Get 2 coolest ====")
	res, err = getNCoolestPresents(append(parray, Present{5, 1}), 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		printSlice(res)
	}

	fmt.Println("\n==== Get -2 coolest ===")
	res, err = getNCoolestPresents(parray, -2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		printSlice(res)
	}

	fmt.Println("\n==== Get 7 coolest ====")
	res, err = getNCoolestPresents(parray, 7)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		printSlice(res)
	}
}
