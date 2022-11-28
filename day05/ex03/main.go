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
	//var grabbed []Present
	if capacity < 0 {
		fmt.Println("error: capacity less than 0")
		return nil
	}

	table := make([][]int, len(ps)+1)
	grabbed := make([][][]Present, len(ps)+1)
	for i := range table {
		table[i] = make([]int, capacity+1)
		grabbed[i] = make([][]Present, capacity+1)
		for j := range grabbed[i] {
			grabbed[i][j] = make([]Present, len(ps))
		}
	}

	for i := 0; i <= len(ps); i++ {
		for j := 0; j <= capacity; j++ {
			if i != 0 && j != 0 {
				fmt.Println("-------- capacity", j, "--------")
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
						//fmt.Println(grabbed[i][j])
						fmt.Println(ps[i-1])
					}
				}
				printTable(table)
			}
			fmt.Println(grabbed[i][j])
		}
	}

	return grabbed[len(ps)][capacity]
}

func main() {
	parray := []Present{{5, 1}, {4, 5}, {5, 2}}
	printHeap(parray)
	grabbed := grabPresents(parray, 5)
	if len(grabbed) != 0 {
		printHeap(grabbed)
	}
}
