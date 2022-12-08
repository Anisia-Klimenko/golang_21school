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

func BenchmarkMinCoins2EmptyInput0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsEmpty[0].val, MinCoins2TestsEmpty[0].coins)
	}
}

func BenchmarkMinCoins2EmptyInput1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsEmpty[1].val, MinCoins2TestsEmpty[2].coins)
	}
}

func BenchmarkMinCoins2EmptyInput2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsEmpty[2].val, MinCoins2TestsEmpty[2].coins)
	}
}

var MinCoins2TestsNegative = []MinCoins2Test{
	{-5, []int{1, 5, 10}, []int{}},
	{10, []int{1, -5, 10}, []int{}},
	{-5, []int{1, -5, 10}, []int{}},
}

func TestMinCoins2NegativeInput(t *testing.T) {
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

func BenchmarkMinCoins2NegativeInput0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsNegative[0].val, MinCoins2TestsNegative[0].coins)
	}
}

func BenchmarkMinCoins2NegativeInput1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsNegative[1].val, MinCoins2TestsNegative[1].coins)
	}
}

func BenchmarkMinCoins2NegativeInput2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsNegative[2].val, MinCoins2TestsNegative[2].coins)
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

func BenchmarkMinCoins2Sorted0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsSorted[0].val, MinCoins2TestsSorted[0].coins)
	}
}

func BenchmarkMinCoins2Sorted1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsSorted[1].val, MinCoins2TestsSorted[1].coins)
	}
}

func BenchmarkMinCoins2Sorted2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsSorted[2].val, MinCoins2TestsSorted[2].coins)
	}
}

func BenchmarkMinCoins2Sorted3(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsSorted[3].val, MinCoins2TestsSorted[3].coins)
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

func BenchmarkMinCoins2Unsorted0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsUnsorted[0].val, MinCoins2TestsUnsorted[0].coins)
	}
}

func BenchmarkMinCoins2Unsorted1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsUnsorted[1].val, MinCoins2TestsUnsorted[1].coins)
	}
}

func BenchmarkMinCoins2Unsorted2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsUnsorted[2].val, MinCoins2TestsUnsorted[2].coins)
	}
}

func BenchmarkMinCoins2Unsorted3(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsUnsorted[3].val, MinCoins2TestsUnsorted[3].coins)
	}
}

func BenchmarkMinCoins2Unsorted4(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsUnsorted[4].val, MinCoins2TestsUnsorted[4].coins)
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

func BenchmarkMinCoins2Imposible0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsImposible[0].val, MinCoins2TestsImposible[0].coins)
	}
}

func BenchmarkMinCoins2Imposible1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsImposible[1].val, MinCoins2TestsImposible[1].coins)
	}
}

func BenchmarkMinCoins2Imposible2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsImposible[2].val, MinCoins2TestsImposible[2].coins)
	}
}

func BenchmarkMinCoins2Imposible3(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsImposible[3].val, MinCoins2TestsImposible[3].coins)
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

func BenchmarkMinCoins2Dublicates0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsDublicates[0].val, MinCoins2TestsDublicates[0].coins)
	}
}

func BenchmarkMinCoins2Dublicates1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsDublicates[1].val, MinCoins2TestsDublicates[1].coins)
	}
}

func BenchmarkMinCoins2Dublicates2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsDublicates[2].val, MinCoins2TestsDublicates[2].coins)
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

func BenchmarkMinCoins2Optimize0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsOptimize[0].val, MinCoins2TestsOptimize[0].coins)
	}
}

func BenchmarkMinCoins2Optimize1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsOptimize[1].val, MinCoins2TestsOptimize[1].coins)
	}
}

func BenchmarkMinCoins2Optimize2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsOptimize[2].val, MinCoins2TestsOptimize[2].coins)
	}
}

func BenchmarkMinCoins2Optimize3(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsOptimize[3].val, MinCoins2TestsOptimize[3].coins)
	}
}
