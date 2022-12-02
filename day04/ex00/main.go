package main

import (
	"errors"
	"log"
	"net/http"
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	//sw "github.com/Anisia-Klimenko/golang_21school/day04/ex00/go-server-server/go"
	//sw "github.com/Anisia-Klimenko/golang_21school/day04/ex00/go"
)

type Candy struct {
	Name  string
	Price int
}

func initCandy() []Candy {
	var menu []Candy
	menu = append(menu, Candy{"CE", 10})
	menu = append(menu, Candy{"AA", 15})
	menu = append(menu, Candy{"NT", 17})
	menu = append(menu, Candy{"DE", 21})
	menu = append(menu, Candy{"YR", 23})
	return menu
}

var menu = initCandy()
var ErrorNotFound = errors.New("candy not found")

func findCandyByName(name string) (Candy, error) {
	for _, candy := range menu {
		if candy.Name == name {
			return candy, nil
		}
	}
	return Candy{}, ErrorNotFound
}

func main() {
	log.Printf("Server started")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":3333", router))
}
