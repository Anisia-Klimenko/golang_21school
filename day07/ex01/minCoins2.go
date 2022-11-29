package ex01

import "sort"

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

func minCoins2(val int, coins []int) []int {
	res1 := countRes(val, coins)
	var res2 []int
	for i := 0; i < len(coins); i++ {
		res2 = countRes(val, coins[:len(coins)-i])
		if len(res2) < len(res1) {
			res1 = res2
		}
	}
	return res1
}
