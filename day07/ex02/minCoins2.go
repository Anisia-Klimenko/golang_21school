/*
	Test MinCoins custom package
*/
package ex02

import "sort"

// MinCoins2 accepts a necessary amount and a sorted slice of unique denominations of coins.
// Updates from the previous version MinCoins:
// 1. checks negative values of coins slice, so it will not have endless loop
// 2. sort coins slice at first, so the algorithm starts getting coins from the biggest denomination
// 3. go through coins slice starting from different coin to find better solution
// 4. checks sum of result coins to fit given value
func MinCoins2(val int, coins []int) []int {
	for _, c := range coins {
		if c <= 0 {
			return []int{}
		}
	}
	res1 := countRes(val, coins)
	var res2 []int
	for i := 0; i < len(coins)/2; i++ {
		res2 = countRes(val, coins[:len(coins)-i])
		if len(res2) < len(res1) {
			res1 = res2
		}
	}
	var sum int
	for _, r := range res1 {
		sum += r
	}
	if sum == val {
		return res1
	} else {
		return []int{}
	}
}

func countRes(val int, coins []int) []int {
	res := make([]int, 0)
	sort.Ints(coins)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}
