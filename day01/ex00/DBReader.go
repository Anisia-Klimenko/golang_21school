package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var DBReader interface {
	read() Cakes
	print(cakes Cakes)
	convert() string
}

type XMLname string

func (filename XMLname) read() Cakes {
	fmt.Println("XMLReader")
	file, err := os.ReadFile(string(filename))
	if err != nil {
		fmt.Println("cannot read input file:", filename)
	}
	_ = file
	var cakes Cakes
	return cakes
}
func (filename XMLname) print(cakes Cakes) {
	fmt.Println("file xml", string(filename))
}

type JSONname string
type Ingredients struct {
	IngredientName  string `json:"ingredient_name"`
	IngredientCount string `json:"ingredient_count"`
	IngredientUnits string `json:"ingredient_unit"`
}
type Cake struct {
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients"`
}
type Cakes struct {
	Cakes []Cake `json:"cake" xml:"cake"`
}

func (filename JSONname) read() Cakes {
	fmt.Println("JSONReader")
	file, err := os.ReadFile(string(filename))
	if err != nil {
		fmt.Println("cannot read input file:", filename)
	}
	var cakes Cakes
	err = json.Unmarshal(file, &cakes)
	//fmt.Println("here", err, cakes.Cakes)

	for i := 0; i < len(cakes.Cakes); i++ {
		fmt.Println("name", cakes.Cakes[i].Name)
		fmt.Println("time", cakes.Cakes[i].Time)
		for j := 0; j < len(cakes.Cakes[i].Ingredients); j++ {
			fmt.Println("\tingredient_name", cakes.Cakes[i].Ingredients[j].IngredientName)
			fmt.Println("\tingredient_count", cakes.Cakes[i].Ingredients[j].IngredientCount)
			fmt.Println("\tingredient_unit", cakes.Cakes[i].Ingredients[j].IngredientUnits, "\n")
		}
	}
	return cakes
}
func (filename JSONname) convert(cakes Cakes) string {
	var result string = "<rec"
	//for cake, i := range cakes.Cakes {
	//
	//}
	return result
}
func (filename JSONname) print(cakes Cakes) {
	for i := 0; i < len(cakes.Cakes); i++ {
		fmt.Println("name", cakes.Cakes[i].Name)
		fmt.Println("time", cakes.Cakes[i].Time)
		for j := 0; j < len(cakes.Cakes[i].Ingredients); j++ {
			fmt.Println("ingredient_name", cakes.Cakes[i].Ingredients[j].IngredientName)
			fmt.Println("ingredient_unit", cakes.Cakes[i].Ingredients[j].IngredientUnits)
		}
	}
}
