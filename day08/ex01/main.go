package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb" json:"name"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(plant interface{}) {
	v := reflect.ValueOf(plant)
	typeS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tag := string(reflect.TypeOf(plant).Field(i).Tag)
		fmt.Printf("%s", typeS.Field(i).Name)
		if tag != "" {
			fmt.Print("(")
			for count, t := range strings.Split(strings.TrimSuffix(tag, " "), " ") {
				if count != 0 {
					fmt.Print(", ")
				}
				split := strings.Split(t, ":")
				fmt.Printf("%s=%s", split[0], strings.Trim(split[1], "\""))
			}
			fmt.Print(")")
		}
		fmt.Printf(": %v\n", v.Field(i).Interface())
	}
}

func main() {
	fmt.Println()
	up1 := UnknownPlant{"rose", "triangle", 12}
	describePlant(up1)

	fmt.Println("----------------------")

	up2 := AnotherUnknownPlant{255, "triangle", 12}
	describePlant(up2)
	fmt.Println()
}
