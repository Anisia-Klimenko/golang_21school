package main

import (
	"flag"
	"fmt"
	"io"
)
import "sort"
import "math"

func mean(bunch []int) float64 {
	var sum float64
	for count, num := range bunch {
		_ = count
		sum += float64(num)
	}
	return math.Round(sum/float64(len(bunch))*100) / 100
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
	return math.Round(math.Pow(sum/float64(len(bunch)), 0.5)*100) / 100
}

var meanf = false
var medianf = false
var modef = false
var sdf = false

func getFlags() {
	flag.BoolVar(&meanf, "mean", false, "mean")
	flag.BoolVar(&medianf, "median", false, "median")
	flag.BoolVar(&modef, "mode", false, "mode")
	flag.BoolVar(&sdf, "sd", false, "sd")
	flag.Parse()

	if !meanf && !medianf && !modef && !sdf {
		meanf = true
		medianf = true
		modef = true
		sdf = true
	}
}

func main() {
	var bunch []int
	var tmp int
	getFlags()
	scanln, err := fmt.Scanln(&tmp)
	if scanln == 0 {
		return
	}
	for err == nil {
		if tmp > 100000 || tmp < -100000 {
			print("number out of range\n")
		} else {
			bunch = append(bunch, tmp)
		}
		scanln, err = fmt.Scanln(&tmp)
	}
	if err != io.EOF {
		fmt.Println(err)
	}
	sort.Sort(sort.IntSlice(bunch))
	if len(bunch) != 0 {
		if meanf {
			fmt.Println("Mean:", mean(bunch))
		}
		if medianf {
			fmt.Println("Median:", median(bunch))
		}
		if modef {
			fmt.Println("Mode:", mode(bunch))
		}
		if sdf {
			fmt.Println("SD:", sd(bunch))
		}
	}
}
