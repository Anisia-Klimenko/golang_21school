package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//func BuyCandy(w http.ResponseWriter, r *http.Request) {
//	decoder := json.NewDecoder(r.Body)
//	var convertJson []byte
//	var order Order
//	err := decoder.Decode(&order)
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		response := InlineResponse400{Error_: "wrong fields or types"}
//		convertJson, err = json.MarshalIndent(response, "", "    ")
//		if err != nil {
//			fmt.Println("Error MarshalIndent:", err)
//		}
//		fmt.Fprintf(w, string(convertJson))
//		return
//	}
//
//	if err != nil {
//		fmt.Println("Error MarshalIndent:", err)
//	}
//	fmt.Fprintf(w, string(convertJson))
//}

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
				log.Println("admin", credentials.AdminLogin, "successfully logged in")
				file, _ := os.ReadFile("resources/adminArticle.html")
				fmt.Fprintf(w, string(file))
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

func generateHTML(article string) (html string) {
	const tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Article Redactor</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">
    <link href="https://fonts.googleapis.com/css2?family=Montserrat:wght@300&display=swap" rel="stylesheet">
</head>
<body style="font-family: 'Montserrat';">
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "Article redactor",
		Items: []string{
			"My photos",
			"My blog",
			article,
		},
	}

	err = t.Execute(os.Stdout, data)
	check(err)

	return
}

func AddArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "adding article...")
	text := r.Form.Get("article")
	generateHTML(text)
	fmt.Fprintf(w, text)
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if !status {
		log.Println("GET admin.html")
		file, _ := os.ReadFile("resources/admin.html")
		fmt.Fprintf(w, string(file))
	} else {
		log.Println("GET adminArticle.html")
		file, _ := os.ReadFile("resources/adminArticle.html")
		fmt.Fprintf(w, string(file))
	}
}
