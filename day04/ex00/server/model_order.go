package main

type Order struct {

	// amount of money put into vending machine
	Money int `json:"money"`

	// kind of candy
	CandyType string `json:"candyType"`

	// number of candy
	CandyCount int `json:"candyCount"`
}

func (order Order) toString() string {
	var result string
	result += "money: " + string(rune(order.Money)) + "\t"
	result += "candyType: " + order.CandyType + "\t"
	result += "candyCount: " + string(rune(order.CandyCount))
	return result
}
