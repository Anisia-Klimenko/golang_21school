package main

type Order struct {

	// amount of money put into vending machine
	Money int `json:"money"`

	// kind of candy
	CandyType string `json:"candyType"`

	// number of candy
	CandyCount int `json:"candyCount"`
}
