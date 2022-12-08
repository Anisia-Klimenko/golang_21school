package ex00

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

type MinCoinsTest struct {
	val      int
	coins    []int
	expected []int
}

var MinCoinsTestsEmpty = []MinCoinsTest{
	{0, []int{1, 5, 10}, []int{}},
	{10, []int{}, []int{}},
	{0, []int{}, []int{}},
}

func TestMinCoinsEmptyInput(t *testing.T) {
	for _, i := range MinCoinsTestsEmpty {
		got := minCoins(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}

var MinCoinsTestsNegative = []MinCoinsTest{
	{-5, []int{1, 5, 10}, []int{}},
	{10, []int{1, -5, 10}, []int{}},
	{-5, []int{1, -5, 10}, []int{}},
}

func TestMinCoinsNegativeInput(t *testing.T) {
	for _, i := range MinCoinsTestsNegative {
		timeout := time.After(3 * time.Second)
		done := make(chan []int)
		go func() {
			got := minCoins(i.val, i.coins)
			done <- got
		}()
		select {
		case <-timeout:
			t.Fatalf("minCoins(%d, %v) -> test didn't finish in time, endless loop", i.val, i.coins)
		case got := <-done:
			sort.Ints(got)
			if !reflect.DeepEqual(got, i.expected) {
				t.Errorf("minCoins(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
			}
		}
	}
}

var MinCoinsTestsSorted = []MinCoinsTest{
	{1, []int{1, 5, 10}, []int{1}},
	{5, []int{1, 5, 10}, []int{5}},
	{10, []int{1, 5, 10}, []int{10}},
	{6, []int{1, 5, 10}, []int{1, 5}},
}

func TestMinCoinsSorted(t *testing.T) {
	for _, i := range MinCoinsTestsSorted {
		got := minCoins(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}

var MinCoinsTestsUnsorted = []MinCoinsTest{
	{6, []int{1, 10, 5}, []int{1, 5}},
	{6, []int{5, 1, 10}, []int{1, 5}},
	{6, []int{5, 10, 1}, []int{1, 5}},
	{6, []int{10, 1, 5}, []int{1, 5}},
	{6, []int{10, 5, 1}, []int{1, 5}},
}

func TestMinCoinsUnsorted(t *testing.T) {
	for _, i := range MinCoinsTestsUnsorted {
		got := minCoins(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}

var MinCoinsTestsImposible = []MinCoinsTest{
	{4, []int{5, 10, 15}, []int{}},
	{9, []int{5, 10, 15}, []int{}},
	{14, []int{5, 10, 15}, []int{}},
	{16, []int{5, 10, 15}, []int{}},
}

func TestMinCoinsImposible(t *testing.T) {
	for _, i := range MinCoinsTestsImposible {
		got := minCoins(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}

var MinCoinsTestsDublicates = []MinCoinsTest{
	{6, []int{1, 1, 5, 10}, []int{1, 5}},
	{6, []int{1, 5, 1, 10}, []int{1, 5}},
	{6, []int{1, 5, 10, 1}, []int{1, 5}},
}

func TestMinCoinsDublicates(t *testing.T) {
	for _, i := range MinCoinsTestsDublicates {
		got := minCoins(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}

var MinCoinsTestsOptimize = []MinCoinsTest{
	{6, []int{4, 1, 3, 3, 4}, []int{3, 3}},
	{15, []int{5, 1, 5, 10}, []int{5, 10}},
	{10, []int{1, 5, 10, 1}, []int{10}},
	{13, []int{5, 10, 1}, []int{1, 1, 1, 10}},
}

func TestMinCoinsOptimize(t *testing.T) {
	for _, i := range MinCoinsTestsOptimize {
		got := minCoins(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}
