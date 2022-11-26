package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

var DBReader interface {
	read() (error, Recipes)
	print(cakes Recipes)
	convert() string
}

type XMLname string
type JSONname string
type Recipes struct {
	Cakes []struct {
		Name        string `json:"name" xml:"name"`
		Time        string `json:"time" xml:"stovetime"`
		Ingredients []struct {
			IngredientName  string `json:"ingredient_name" xml:"item>itemname"`
			IngredientCount string `json:"ingredient_count" xml:"item>itemcount"`
			IngredientUnit  string `json:"ingredient_unit,omitempty" xml:"item>itemunit"`
		} `json:"ingredients" xml:"ingredients"`
	} `json:"cake" xml:"cake"`
}

func (filename XMLname) read() (error, Recipes) {
	fmt.Println("XML DBReader")
	file, err := os.ReadFile(string(filename))
	if err != nil {
		fmt.Println("cannot read input file:", filename, err)
		return err, Recipes{nil}
	}
	_ = file
	var cakes Recipes
	err = xml.Unmarshal(file, &cakes)
	if err != nil {
		fmt.Println("can't read parse file", err)
	}
	return err, cakes
}
func (filename XMLname) convert(cakes Recipes) string {
	b, err := json.MarshalIndent(cakes, "", "    ")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
func (filename XMLname) print(cakes Recipes) {
	for i := 0; i < len(cakes.Cakes); i++ {
		fmt.Println("name", cakes.Cakes[i].Name)
		fmt.Println("time", cakes.Cakes[i].Time)
		for j := 0; j < len(cakes.Cakes[i].Ingredients); j++ {
			fmt.Println("\tingredient_name", cakes.Cakes[i].Ingredients[j].IngredientName)
			fmt.Println("\tingredient_count", cakes.Cakes[i].Ingredients[j].IngredientCount)
			fmt.Println("\tingredient_unit", cakes.Cakes[i].Ingredients[j].IngredientUnit, "\n")
		}
	}
}

func (filename JSONname) read() (error, Recipes) {
	fmt.Println("JSON DBReader")
	file, err := os.ReadFile(string(filename))
	if err != nil {
		fmt.Println("cannot read input file:", filename)
		return err, Recipes{}
	}
	var cakes Recipes
	err = json.Unmarshal(file, &cakes)
	if err != nil {
		fmt.Println("can't read parse file", err)
	}
	return err, cakes
}
func (filename JSONname) convert(cakes Recipes) string {
	b, err := xml.MarshalIndent(cakes, "", "    ")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
func (filename JSONname) print(cakes Recipes) {
	for i := 0; i < len(cakes.Cakes); i++ {
		fmt.Println("name", cakes.Cakes[i].Name)
		fmt.Println("time", cakes.Cakes[i].Time)
		for j := 0; j < len(cakes.Cakes[i].Ingredients); j++ {
			fmt.Println("\tingredient_name", cakes.Cakes[i].Ingredients[j].IngredientName)
			fmt.Println("\tingredient_count", cakes.Cakes[i].Ingredients[j].IngredientCount)
			fmt.Println("\tingredient_unit", cakes.Cakes[i].Ingredients[j].IngredientUnit, "\n")
		}
	}
}
