package ex00

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

type MinCoins2Test struct {
	val      int
	coins    []int
	expected []int
}

var MinCoins2TestsEmpty = []MinCoins2Test{
	{0, []int{1, 5, 10}, []int{}},
	{10, []int{}, []int{}},
	{0, []int{}, []int{}},
}

func TestMinCoins2EmptyInput(t *testing.T) {
	for _, i := range MinCoins2TestsEmpty {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins2(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}

var MinCoins2TestsNegative = []MinCoins2Test{
	{-5, []int{1, 5, 10}, []int{}},
	{10, []int{1, -5, 10}, []int{}},
	{-5, []int{1, -5, 10}, []int{}},
}

func TestMinCoins2NegativeInput(t *testing.T) {
	//go func() {
	//	// do your testing
	//	time.Sleep(5 * time.Second)
	//	done <- true
	//}()
	//
	//select {
	//case <-timeout:
	//	t.Fatal("Test didn't finish in time")
	//case <-done:
	//}
	for _, i := range MinCoins2TestsNegative {
		timeout := time.After(3 * time.Second)
		done := make(chan []int)
		go func() {
			got := minCoins2(i.val, i.coins)
			done <- got
		}()
		select {
		case <-timeout:
			t.Fatal("Test didn't finish in time")
		case got := <-done:
			sort.Ints(got)
			if !reflect.DeepEqual(got, i.expected) {
				t.Errorf("minCoins2(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
			}
		}
	}
}

var MinCoins2TestsSorted = []MinCoins2Test{
	{1, []int{1, 5, 10}, []int{1}},
	{5, []int{1, 5, 10}, []int{5}},
	{10, []int{1, 5, 10}, []int{10}},
	{6, []int{1, 5, 10}, []int{1, 5}},
}

func TestMinCoins2Sorted(t *testing.T) {
	for _, i := range MinCoins2TestsSorted {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins2(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}

var MinCoins2TestsUnsorted = []MinCoins2Test{
	{6, []int{1, 10, 5}, []int{1, 5}},
	{6, []int{5, 1, 10}, []int{1, 5}},
	{6, []int{5, 10, 1}, []int{1, 5}},
	{6, []int{10, 1, 5}, []int{1, 5}},
	{6, []int{10, 5, 1}, []int{1, 5}},
}

func TestMinCoins2Unsorted(t *testing.T) {
	for _, i := range MinCoins2TestsUnsorted {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins2(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}

var MinCoins2TestsImposible = []MinCoins2Test{
	{4, []int{5, 10, 15}, []int{}},
	{9, []int{5, 10, 15}, []int{}},
	{14, []int{5, 10, 15}, []int{}},
	{16, []int{5, 10, 15}, []int{}},
}

func TestMinCoins2Imposible(t *testing.T) {
	for _, i := range MinCoins2TestsImposible {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins2(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}

var MinCoins2TestsDublicates = []MinCoins2Test{
	{6, []int{1, 1, 5, 10}, []int{1, 5}},
	{6, []int{1, 5, 1, 10}, []int{1, 5}},
	{6, []int{1, 5, 10, 1}, []int{1, 5}},
}

func TestMinCoins2Dublicates(t *testing.T) {
	for _, i := range MinCoins2TestsDublicates {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins2(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}

var MinCoins2TestsOptimize = []MinCoins2Test{
	{6, []int{4, 1, 3, 3}, []int{3, 3}},
	{15, []int{5, 1, 5, 10}, []int{5, 10}},
	{10, []int{1, 5, 10, 1}, []int{10}},
	{13, []int{5, 10, 1}, []int{1, 1, 1, 10}},
}

func TestMinCoins2Optimize(t *testing.T) {
	for _, i := range MinCoins2TestsOptimize {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.expected) {
			t.Errorf("minCoins2(%d, %v) = %v, expected %v", i.val, i.coins, got, i.expected)
		}
	}
}
