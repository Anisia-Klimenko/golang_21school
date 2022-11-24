package main

import "fmt"
import "sort"
import "math"

func mean(bunch []int) float64 {
	var sum float64
	for count, num := range bunch {
		_ = count
		sum += float64(num)
	}
	return sum / float64(len(bunch))
}

func median(bunch []int) float64 {
	if len(bunch)%2 == 1 {
		return float64(bunch[len(bunch)/2])
	} else {
		return mean([]int{bunch[len(bunch)/2-1], bunch[len(bunch)/2]})
	}
}

func mode(bunch []int) int {
	var curFreqElem int
	var curFreq int
	var freqElem int = bunch[0]
	var freq int = 1
	for count, num := range bunch {
		_ = num
		if count > 0 {
			if bunch[count] == bunch[count-1] {
				curFreq++
				curFreqElem = bunch[count]
			} else {
				if curFreq > freq {
					freq = curFreq
					freqElem = curFreqElem
				}
				curFreq = 1
			}
		}
	}
	return freqElem
}

func sd(bunch []int) float64 {
	var mean = mean(bunch)
	var sum float64
	for count, num := range bunch {
		_ = count
		cur := float64(num)
		sum += math.Pow(cur-mean, 2)
	}
	return math.Pow(sum/float64(len(bunch)), 0.5)
}

func main() {
	var bunch []int
	var tmp int
	scanln, err := fmt.Scanln(&tmp)
	_ = scanln
	for err == nil {
		bunch = append(bunch, tmp)
		scanln, err = fmt.Scanln(&tmp)
	}
	sort.Sort(sort.IntSlice(bunch))
	fmt.Println("Mean:", mean(bunch))
	fmt.Println("Median:", median(bunch))
	fmt.Println("Mode:", mode(bunch))
	fmt.Println("SD:", sd(bunch))
}
