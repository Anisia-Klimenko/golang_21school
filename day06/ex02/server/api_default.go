package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !status {
		if r.Form.Get("login") == credentials.AdminLogin {
			if r.Form.Get("password") == credentials.AdminPassword {
				status = true
				log.Println("admin", credentials.AdminLogin, "successfully logged in")
				AllPosts(w, r)
			} else {
				status = false
				log.Println("admin", credentials.AdminLogin, "log in error, wrong password")
				fmt.Fprintf(w, "bad password")
			}
		} else {
			status = false
			log.Println("admin", credentials.AdminLogin, "log in error, no such login")
			fmt.Fprintf(w, "no such admin")
		}
	} else {
		AddArticle(w, r)
	}
}

func AddArticle(w http.ResponseWriter, r *http.Request) {
	InsertPost(w, r)
	AllPosts(w, r)
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if !status {
		log.Println("GET admin.html")
		file, _ := os.ReadFile("resources/admin.html")
		fmt.Fprintf(w, string(file))
	} else {
		AllPosts(w, r)
	}
}
