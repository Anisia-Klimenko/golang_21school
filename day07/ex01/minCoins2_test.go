package ex01

import (
	"reflect"
	"sort"
	"testing"
)

var vals = []int{0, 10, 0, -10, 13, 13, 13, 10, 13, 15, 10, 15, 6}
var coins = [][]int{
	{1, 5, 10},
	{},
	{},
	{1, 5, 10},
	{5, 1, 10},
	{5, 10, 1},
	{10, 5, 1},
	{1, 10, 5},
	{1, 5, 10},
	{1, 5, 10},
	{1, 5, 10},
	{5, 1, 5, 10},
	{4, 1, 3, 3}}
var res = [][]int{{}, {}, {}, {}, {1, 1, 1, 10}, {1, 1, 1, 10}, {1, 1, 1, 10}, {10}, {1, 1, 1, 10}, {5, 10}, {10}, {5, 10}, {3, 3}}

func TestMinCoins(t *testing.T) {
	for i := range coins {
		got := minCoins2(vals[i], coins[i])
		sort.Ints(got)
		if !reflect.DeepEqual(got, res[i]) {
			t.Errorf("minCoins(%d, %v) = %v, expected %v", vals[i], coins[i], got, res[i])
		}
	}
}
