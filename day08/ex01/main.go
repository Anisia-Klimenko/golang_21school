package main

import (
	"fmt"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(plant interface{}) {
	//typeOf := reflect.TypeOf(plant)
	//tag := string(field.Tag)

	v := reflect.ValueOf(plant)
	typeS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tag := reflect.TypeOf(plant).Field(i).Tag
		fmt.Printf("%s", typeS.Field(i).Name)
		if tag != "" {
			fmt.Printf("(%s)", tag)
		}
		fmt.Printf(": %v\n", v.Field(i).Interface())
	}
}

func main() {
	up1 := UnknownPlant{"rose", "triangle", 12}
	describePlant(up1)
}
