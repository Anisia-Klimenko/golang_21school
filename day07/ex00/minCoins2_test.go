package ex00

import (
	"reflect"
	"sort"
	"testing"
)

func TestMinCoins2(t *testing.T) {
	for i := range coins {
		got := minCoins2(vals[i], coins[i])
		sort.Ints(got)
		if !reflect.DeepEqual(got, res[i]) {
			t.Errorf("minCoins(%d, %v) = %v, expected %v", vals[i], coins[i], got, res[i])
		}
	}
}

func TestMinCoins2EmptyVal(t *testing.T) {
	coins := []int{1, 5, 10}
	val := 0
	got := minCoins2(val, coins)
	var res []int
	if len(got) != 0 {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2EmptyCoins(t *testing.T) {
	var coins []int
	val := 10
	got := minCoins2(val, coins)
	var res []int
	if len(got) != 0 {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2EmptyCoinsAndVal(t *testing.T) {
	var coins []int
	val := 0
	got := minCoins2(val, coins)
	var res []int
	if len(got) != 0 {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2NegativeVal(t *testing.T) {
	coins := []int{1, 5, 10}
	val := -10
	got := minCoins2(val, coins)
	var res []int
	if len(got) != 0 {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2UnsortedCoins1(t *testing.T) {
	coins := []int{5, 1, 10}
	val := 13
	got := minCoins2(val, coins)
	sort.Ints(got)
	res := []int{1, 1, 1, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2UnsortedCoins2(t *testing.T) {
	coins := []int{5, 10, 1}
	val := 13
	got := minCoins2(val, coins)
	sort.Ints(got)
	res := []int{1, 1, 1, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2UnsortedCoins3(t *testing.T) {
	coins := []int{10, 5, 1}
	val := 13
	got := minCoins2(val, coins)
	sort.Ints(got)
	res := []int{1, 1, 1, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}
func TestMinCoins2UnsortedCoins4(t *testing.T) {
	coins := []int{1, 10, 5}
	val := 10
	got := minCoins2(val, coins)
	sort.Ints(got)
	res := []int{10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2SortedCoins1(t *testing.T) {
	coins := []int{1, 5, 10}
	val := 13
	got := minCoins2(val, coins)
	sort.Ints(got)
	res := []int{1, 1, 1, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2SortedCoins2(t *testing.T) {
	coins := []int{1, 5, 10}
	val := 15
	got := minCoins2(val, coins)
	sort.Ints(got)
	res := []int{5, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2SortedCoins3(t *testing.T) {
	coins := []int{1, 5, 10}
	val := 10
	got := minCoins2(val, coins)
	sort.Ints(got)
	res := []int{10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2RecurringCoins(t *testing.T) {
	coins := []int{5, 1, 5, 10}
	val := 15
	got := minCoins2(val, coins)
	sort.Ints(got)
	res := []int{5, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins2(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoins2RecurringCoins2(t *testing.T) {
	coins := []int{4, 1, 3, 3}
	val := 6
	got := minCoins2(val, coins)
	sort.Ints(got)
	res := []int{3, 3}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}
