package main

import (
	"log"
	"net/http"
	"os"
	"strings"
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

type Credentials struct {
	AdminLogin    string
	AdminPassword string
	DBLogin       string
	DBPassword    string
}

var credentials Credentials
var status = false

func initCredentials() {
	file, _ := os.ReadFile("../admin_credentials.txt")
	for _, line := range strings.Split(strings.TrimSuffix(string(file), "\n"), "\n") {
		if strings.Contains(line, "admin_login") {
			credentials.AdminLogin = strings.Split(line, "=")[1]
		} else if strings.Contains(line, "admin_password") {
			credentials.AdminPassword = strings.Split(line, "=")[1]
		} else if strings.Contains(line, "db_login") {
			credentials.DBLogin = strings.Split(line, "=")[1]
		} else if strings.Contains(line, "db_password") {
			credentials.DBPassword = strings.Split(line, "=")[1]
		}
	}
}

func main() {
	log.Printf("Server started")
	initCredentials()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}
