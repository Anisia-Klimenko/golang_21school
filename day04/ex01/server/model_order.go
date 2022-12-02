package main

type Order struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

func (order Order) toString() string {
	var result string
	result += "money: " + string(rune(order.Money)) + "\t"
	result += "candyType: " + order.CandyType + "\t"
	result += "candyCount: " + string(rune(order.CandyCount))
	return result
}
