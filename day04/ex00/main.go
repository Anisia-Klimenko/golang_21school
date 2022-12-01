package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "github.com/Anisia-Klimenko/golang_21school/day04/ex00/go-server-server/go"
	//sw "github.com/Anisia-Klimenko/golang_21school/day04/ex00/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":3333", router))
}
