package ex00

import (
	"reflect"
	"sort"
	"testing"
)

func TestMinCoinsEmptyVal(t *testing.T) {
	coins := []int{1, 5, 10}
	val := 0
	got := minCoins(val, coins)
	var res []int
	if len(got) != 0 {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsEmptyCoins(t *testing.T) {
	var coins []int
	val := 10
	got := minCoins(val, coins)
	var res []int
	if len(got) != 0 {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsEmptyCoinsAndVal(t *testing.T) {
	var coins []int
	val := 0
	got := minCoins(val, coins)
	var res []int
	if len(got) != 0 {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsNegativeVal(t *testing.T) {
	coins := []int{1, 5, 10}
	val := -10
	got := minCoins(val, coins)
	var res []int
	if len(got) != 0 {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsUnsortedCoins1(t *testing.T) {
	coins := []int{5, 1, 10}
	val := 13
	got := minCoins(val, coins)
	sort.Ints(got)
	res := []int{1, 1, 1, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsUnsortedCoins2(t *testing.T) {
	coins := []int{5, 10, 1}
	val := 13
	got := minCoins(val, coins)
	sort.Ints(got)
	res := []int{1, 1, 1, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsUnsortedCoins3(t *testing.T) {
	coins := []int{10, 5, 1}
	val := 13
	got := minCoins(val, coins)
	sort.Ints(got)
	res := []int{1, 1, 1, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}
func TestMinCoinsUnsortedCoins4(t *testing.T) {
	coins := []int{1, 10, 5}
	val := 10
	got := minCoins(val, coins)
	sort.Ints(got)
	res := []int{10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsSortedCoins1(t *testing.T) {
	coins := []int{1, 5, 10}
	val := 13
	got := minCoins(val, coins)
	sort.Ints(got)
	res := []int{1, 1, 1, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsSortedCoins2(t *testing.T) {
	coins := []int{1, 5, 10}
	val := 15
	got := minCoins(val, coins)
	sort.Ints(got)
	res := []int{5, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsSortedCoins3(t *testing.T) {
	coins := []int{1, 5, 10}
	val := 10
	got := minCoins(val, coins)
	sort.Ints(got)
	res := []int{10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsRecurringCoins(t *testing.T) {
	coins := []int{5, 1, 5, 10}
	val := 15
	got := minCoins(val, coins)
	sort.Ints(got)
	res := []int{5, 10}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}

func TestMinCoinsRecurringCoins2(t *testing.T) {
	coins := []int{4, 1, 3, 3}
	val := 6
	got := minCoins(val, coins)
	sort.Ints(got)
	res := []int{3, 3}
	if !reflect.DeepEqual(got, res) {
		t.Errorf("minCoins(%d, %v) = %v, expected %v", val, coins, got, res)
	}
}
