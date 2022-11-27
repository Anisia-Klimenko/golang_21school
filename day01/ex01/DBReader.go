package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"reflect"
)

type DBReader interface {
	read() (error, Recipes)
	print(cakes Recipes)
	convert(cakes Recipes) string
}

type XMLname string
type JSONname string
type Recipes struct {
	Cakes []Cakes `json:"cake" xml:"cake"`
}
type Cakes struct {
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
}
type Ingredients struct {
	IngredientName  string `json:"ingredient_name" xml:"itemname"`
	IngredientCount string `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  string `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

func (filename XMLname) read() (error, Recipes) {
	file, err := os.ReadFile(string(filename))
	if err != nil {
		fmt.Println("cannot read input file:", filename, err)
		return err, Recipes{nil}
	}
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

func nameExists(cakes Recipes, name string) bool {
	if getCakeIndexByName(cakes, name) >= 0 {
		return true
	}
	return false
}

func equalTime(cakes Recipes, time string) bool {
	for _, s := range cakes.Cakes {
		if time == s.Time {
			return true
		}
	}
	return false
}

func getCakeIndexByName(cakes Recipes, name string) int {
	for count, s := range cakes.Cakes {
		if name == s.Name {
			return count
		}
	}
	return -1
}

func ingredientExists(cake Cakes, ingredient string) bool {
	for _, s := range cake.Ingredients {
		if ingredient == s.IngredientName {
			return true
		}
	}
	return false
}

// 0 - not equal
// 1 - in cake ingredient unit is missing
// 2 - in ingredient unit is missing
// 3 - equal
func equalIngredientUnit(cake Cakes, ingredient Ingredients) int {
	for _, s := range cake.Ingredients {
		if ingredient.IngredientName == s.IngredientName {
			if ingredient.IngredientUnit == s.IngredientUnit {
				return 3
			} else if len(s.IngredientUnit) == 0 {
				return 1
			} else if len(ingredient.IngredientUnit) == 0 {
				return 2
			} else {
				return 0
			}
		}
	}
	return 3
}

func getIngredientIndexByName(ingredients []Ingredients, name string) int {
	for count, s := range ingredients {
		if name == s.IngredientName {
			return count
		}
	}
	return -1
}

func equalIngredientCount(ingredients []Ingredients, count string) bool {
	for _, s := range ingredients {
		if count == s.IngredientCount {
			return true
		}
	}
	return false
}

func compare(old Recipes, new Recipes) {
	if reflect.DeepEqual(old, new) {
		fmt.Println("EQUAL")
		return
	}

	// added/removed cake
	for i := 0; i < len(new.Cakes); i++ {
		if !nameExists(old, new.Cakes[i].Name) {
			fmt.Printf("%sADDED%s cake \"%s\"\n", GRN, END, new.Cakes[i].Name)
		}
	}
	for i := 0; i < len(old.Cakes); i++ {
		if !nameExists(new, old.Cakes[i].Name) {
			fmt.Printf("%sREMOVED%s cake \"%s\"\n", RED, END, new.Cakes[i].Name)
		}
	}

	// changed time
	for i := 0; i < len(new.Cakes); i++ {
		if nameExists(old, new.Cakes[i].Name) {
			if !equalTime(old, new.Cakes[i].Time) {
				fmt.Printf("%sCHANGED%s cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", YELLOW, END,
					new.Cakes[i].Name, new.Cakes[i].Time, old.Cakes[getCakeIndexByName(old, new.Cakes[i].Name)].Time)
			}
		}
	}

	// added/removed ingredient
	for i := 0; i < len(new.Cakes); i++ {
		if nameExists(old, new.Cakes[i].Name) {
			for j := 0; j < len(new.Cakes[i].Ingredients); j++ {
				if !ingredientExists(old.Cakes[getCakeIndexByName(old, new.Cakes[i].Name)], new.Cakes[i].Ingredients[j].IngredientName) {
					fmt.Printf("%sADDED%s ingredient \"%s\" for cake  \"%s\"\n", GRN, END, new.Cakes[i].Ingredients[j].IngredientName, new.Cakes[i].Name)
				}
			}
		}
	}
	for i := 0; i < len(old.Cakes); i++ {
		if nameExists(new, old.Cakes[i].Name) {
			for j := 0; j < len(old.Cakes[i].Ingredients); j++ {
				if !ingredientExists(new.Cakes[getCakeIndexByName(new, old.Cakes[i].Name)], old.Cakes[i].Ingredients[j].IngredientName) {
					fmt.Printf("%sREMOVED%s ingredient \"%s\" for cake  \"%s\"\n", RED, END, old.Cakes[i].Ingredients[j].IngredientName, old.Cakes[i].Name)
				}
			}
		}
	}

	// changed ingredient count
	for i := 0; i < len(new.Cakes); i++ {
		if nameExists(old, new.Cakes[i].Name) {
			for j := 0; j < len(new.Cakes[i].Ingredients); j++ {
				oldCake := old.Cakes[getCakeIndexByName(new, old.Cakes[i].Name)]
				if !equalIngredientCount(oldCake.Ingredients, new.Cakes[i].Ingredients[j].IngredientCount) {
					fmt.Printf("%sCHANGED%s unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", YELLOW, END,
						new.Cakes[i].Ingredients[j].IngredientName, new.Cakes[i].Name, new.Cakes[i].Ingredients[j].IngredientCount,
						oldCake.Ingredients[getIngredientIndexByName(oldCake.Ingredients, new.Cakes[i].Ingredients[j].IngredientName)].IngredientCount)
				}
			}
		}
	}

	// changed ingredient unit
	for i := 0; i < len(new.Cakes); i++ {
		if nameExists(old, new.Cakes[i].Name) {
			for j := 0; j < len(new.Cakes[i].Ingredients); j++ {
				oldCake := old.Cakes[getCakeIndexByName(new, old.Cakes[i].Name)]
				if equalIngredientUnit(oldCake, new.Cakes[i].Ingredients[j]) == 0 {
					fmt.Printf("%sCHANGED%s unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", YELLOW, END,
						new.Cakes[i].Ingredients[j].IngredientName, new.Cakes[i].Name, new.Cakes[i].Ingredients[j].IngredientUnit,
						oldCake.Ingredients[getIngredientIndexByName(oldCake.Ingredients, new.Cakes[i].Ingredients[j].IngredientName)].IngredientUnit)
				} else if equalIngredientUnit(oldCake, new.Cakes[i].Ingredients[j]) == 1 {
					fmt.Printf("%sADDED%s unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", GRN, END,
						new.Cakes[i].Ingredients[j].IngredientUnit, new.Cakes[i].Ingredients[j].IngredientName, new.Cakes[i].Name)
				} else if equalIngredientUnit(oldCake, new.Cakes[i].Ingredients[j]) == 2 {
					fmt.Printf("%sREMOVED%s unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", RED, END,
						oldCake.Ingredients[getIngredientIndexByName(oldCake.Ingredients, new.Cakes[i].Ingredients[j].IngredientName)].IngredientUnit,
						new.Cakes[i].Ingredients[j].IngredientName, new.Cakes[i].Name)
				}
			}
		}
	}
}
