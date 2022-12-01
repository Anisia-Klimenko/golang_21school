package swagger

type Order struct {

	// amount of money put into vending machine
	Money int32 `json:"money"`

	// kind of candy
	CandyType string `json:"candyType"`

	// number of candy
	CandyCount int32 `json:"candyCount"`
}

func (order Order) toString() string {
	var result string
	result += "money: " + string(order.Money) + "\t"
	result += "candyType: " + order.CandyType + "\t"
	result += "candyCount: " + string(order.CandyCount)
	return result
}
