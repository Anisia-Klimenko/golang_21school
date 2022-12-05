package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var convertJson []byte
	var order Order
	err := decoder.Decode(&order)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := InlineResponse400{Error_: "wrong fields or types"}
		convertJson, err = json.MarshalIndent(response, "", "    ")
		if err != nil {
			fmt.Println("Error MarshalIndent:", err)
		}
		fmt.Fprintf(w, string(convertJson))
		return
	}

	if err != nil {
		fmt.Println("Error MarshalIndent:", err)
	}
	fmt.Fprintf(w, string(convertJson))
}

type AdmCred struct {
	Login  string `json:"login"`
	Passwd string `json:"password"`
}

func Admin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !status {
		if r.Form.Get("login") == credentials.AdminLogin {
			if r.Form.Get("password") == credentials.AdminPassword {
				status = true
				file, _ := os.ReadFile("../adminArticle.html")
				fmt.Fprintf(w, string(file))
			} else {
				status = false
				fmt.Fprintf(w, "bad password")
			}
		} else {
			status = false
			fmt.Fprintf(w, "no such admin")
		}
	} else {
		AddArticle(w, r)
	}
}

func AddArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "adding article...")
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if !status {
		file, _ := os.ReadFile("../admin.html")
		fmt.Fprintf(w, string(file))
	} else {
		file, _ := os.ReadFile("../adminArticle.html")
		fmt.Fprintf(w, string(file))
	}
}
