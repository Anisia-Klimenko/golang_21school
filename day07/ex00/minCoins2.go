package ex00

import "sort"

func minCoins2(val int, coins []int) []int {
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
