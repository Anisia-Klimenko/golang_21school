package main

import (
	"flag"
	"fmt"
	"path"
)

func main() {
	var fileName *string
	fileName = flag.String("f", "", "input file name")
	flag.Parse()
	fmt.Println("file name", *fileName)
	if path.Ext(*fileName) == ".xml" {
		content := XMLname(*fileName).read()

		//_ = content
		XMLname(*fileName).print(content)
	} else if path.Ext(*fileName) == ".json" {
		content := JSONname(*fileName).read()
		//JSONname(content).print()
		_ = content
	} else {
		fmt.Println("usage: ./readDB -f <filename>.xml")
		fmt.Println("       ./readDB -f <filename>.json")
	}
}

//func main() {
//	var jsonBlob = []byte(`{
//  "cake": [
//    {
//      "name": "Red Velvet Strawberry Cake",
//      "time": "45 min",
//      "ingredients": [
//        {
//          "ingredient_name": "Flour",
//          "ingredient_count": "2",
//          "ingredient_unit": "mugs"
//        },
//        {
//          "ingredient_name": "Strawberries",
//          "ingredient_count": "7"
//        },
//        {
//          "ingredient_name": "Vanilla extract",
//          "ingredient_count": "2.5",
//          "ingredient_unit": "tablespoons"
//        }
//      ]
//    },
//    {
//      "name": "Blueberry Muffin Cake",
//      "time": "30 min",
//      "ingredients": [
//        {
//          "ingredient_name": "Brown sugar",
//          "ingredient_count": "1",
//          "ingredient_unit": "mug"
//        },
//        {
//          "ingredient_name": "Blueberries",
//          "ingredient_count": "1",
//          "ingredient_unit": "mug"
//        }
//      ]
//    }
//  ]
//}`)
//	type Ingredients struct {
//		IngredientName  string `json:"ingredient_name"`
//		IngredientCount string `json:"ingredient_count"`
//		IngredientUnits string `json:"ingredient_units"`
//	}
//	type Cake struct {
//		Name        string        `json:"name"`
//		Time        string        `json:"time"`
//		Ingredients []Ingredients `json:"ingredients"`
//	}
//	type Cakes struct {
//		Cakes []Cake `json:"cake"`
//	}
//
//	var cakes Cakes
//	err := json.Unmarshal(jsonBlob, &cakes)
//	_ = err
//	//fmt.Println("here", err, cakes.Cakes)
//
//	fmt.Println(len(cakes.Cakes), err)
//	for i := 0; i < len(cakes.Cakes); i++ {
//		fmt.Println("name", cakes.Cakes[i].Name)
//		fmt.Println("time", cakes.Cakes[i].Time)
//		for j := 0; j < len(cakes.Cakes[i].Ingredients); j++ {
//			fmt.Println("ingredient_name", cakes.Cakes[i].Ingredients[j].IngredientName)
//		}
//	}
//}
