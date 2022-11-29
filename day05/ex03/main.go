package main

import (
	"fmt"
	"math"
)

type Present struct {
	Value int
	Size  int
}

func printHeap(p []Present) {
	for count, pr := range p {
		fmt.Println(count, " -> value:", pr.Value, "size:", pr.Size)
	}
}

func printTable(table [][]int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			fmt.Print("\t", table[i][j])

		}
		fmt.Println()
	}
}

func appendPresent(ps []Present, p Present) []Present {
	if p.Value == 0 && p.Size == 0 {
		return ps
	}
	for _, present := range ps {
		if present.Value == p.Value && present.Size == p.Size {
			return ps
		}
	}
	return append(ps, p)
}

func grabPresents(ps []Present, capacity int) []Present {
	if capacity < 0 {
		fmt.Println("error: capacity less than 0")
		return nil
	}
	//if len(ps) == 0 {
	//	return ps
	//}

	table := make([][]int, len(ps)+1)
	grabbed := make([][][]Present, len(ps)+1)
	for i := range table {
		table[i] = make([]int, capacity+1)
		grabbed[i] = make([][]Present, capacity+1)
	}
	//printTable(table)

	for i := 0; i <= len(ps); i++ {
		for j := 0; j <= capacity; j++ {
			if i != 0 && j != 0 {
				//fmt.Println("-------- capacity", j, "--------")
				if ps[i-1].Size > j {
					table[i][j] = table[i-1][j]
					grabbed[i][j] = grabbed[i-1][j]
				} else {
					var prev = table[i-1][j]
					var result = ps[i-1].Value + table[i-1][j-ps[i-1].Size]
					table[i][j] = int(math.Max(float64(prev), float64(result)))
					if table[i][j] == prev {
						grabbed[i][j] = grabbed[i-1][j]
					} else {
						grabbed[i][j] = appendPresent(grabbed[i-1][j-ps[i-1].Size], ps[i-1])
						//fmt.Println(ps[i-1])
					}
				}
				//printTable(table)
			}
			//fmt.Println(grabbed[i][j])
		}
	}
	return grabbed[len(ps)][capacity]
}

func main() {
	fmt.Println("---------------- all presents ----------------")
	parray := []Present{{5, 1}, {4, 5}, {5, 2}}
	printHeap(parray)
	fmt.Println("-------- grabbed presents, capacity 5 --------")
	grabbed := grabPresents(parray, 5)
	if len(grabbed) != 0 {
		printHeap(grabbed)
	}
	fmt.Println()

	fmt.Println("---------------- all presents ----------------")
	parray = []Present{{4000, 4}, {2500, 1}, {2000, 3}}
	printHeap(parray)
	fmt.Println("-------- grabbed presents, capacity 4 --------")
	grabbed = grabPresents(parray, 4)
	if len(grabbed) != 0 {
		printHeap(grabbed)
	}
	fmt.Println()

	fmt.Println("---------------- all presents ----------------")
	parray = []Present{{4000, 4}, {2500, 1}, {2000, 3}}
	printHeap(parray)
	fmt.Println("-------- grabbed presents, capacity 0 --------")
	grabbed = grabPresents(parray, 0)
	if len(grabbed) != 0 {
		printHeap(grabbed)
	}
	fmt.Println("-------- grabbed presents, capacity -1 -------")
	grabbed = grabPresents(parray, -1)
	if len(grabbed) != 0 {
		printHeap(grabbed)
	}
	fmt.Println()

	fmt.Println("---------------- all presents ----------------")
	parray = []Present{{5, 1}, {4, 5}, {5, 2}, {4, 6}, {4, 1},
		{4, 2}, {3, 1}, {5, 6}}
	printHeap(parray)
	fmt.Println("-------- grabbed presents, capacity 10 -------")
	grabbed = grabPresents(parray, 10)
	if len(grabbed) != 0 {
		printHeap(grabbed)
	}
	fmt.Println()

	fmt.Println("---------------- all presents ----------------")
	parray = []Present{}
	printHeap(parray)
	fmt.Println("-------- grabbed presents, capacity 10 -------")
	grabbed = grabPresents(parray, 10)
	if len(grabbed) != 0 {
		printHeap(grabbed)
	}
	fmt.Println()

	fmt.Println("---------------- all presents ----------------")
	parray = []Present{{5, 10}, {4, 50}, {5, 20}, {4, 60}, {4, 10},
		{4, 20}, {3, 10}, {5, 60}}
	printHeap(parray)
	fmt.Println("-------- grabbed presents, capacity 1 --------")
	grabbed = grabPresents(parray, 1)
	if len(grabbed) != 0 {
		printHeap(grabbed)
	}
	fmt.Println()

	fmt.Println("---------------- all presents ----------------")
	parray = []Present{{5, 5}, {5, 5}, {5, 5}}
	printHeap(parray)
	fmt.Println("-------- grabbed presents, capacity 5 --------")
	grabbed = grabPresents(parray, 5)
	if len(grabbed) != 0 {
		printHeap(grabbed)
	}
	fmt.Println()
}
