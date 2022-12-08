package ex02

import (
	"fmt"
)

func ExampleMinCoins2() {
	result := MinCoins2(6, []int{4, 1, 3, 3})
	fmt.Println("MinCoins2(6, []int{4, 1, 3, 3}) = ", result)

	// Output:
	// MinCoins2(6, []int{4, 1, 3, 3}) =  [3 3]
}
