package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	_ "github.com/go-pg/pg/orm"
	"github.com/gorilla/context"
	"html/template"
	"log"
	"net/http"
)

type Article struct {
	Id      int64  `pg:"articles.id"`
	Preview string `pg:"articles.preview"`
	Article string `pg:"articles.article"`
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Article)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	db := DBConn()
	createSchema(db)
}

func DBConn() (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		Database: "postgres",
		User:     credentials.DBLogin,
		Password: credentials.DBPassword,
	})
	return db
}

// NewPost Return new blog Post html form on GET
func NewPost(w http.ResponseWriter, r *http.Request) {
	data := context.Get(r, "article")
	t, _ := template.ParseFiles("templates/layout.html", "templates/new.html")
	t.Execute(w, data)
	http.Redirect(w, r, "/admin", 301)
}

func NewGet(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/layout.html", "templates/new.html")
	t.Execute(w, nil)
}

// AllPosts Get all blog posts and render template
func AllPosts(w http.ResponseWriter, r *http.Request) {
	db := DBConn()
	var posts []Article
	err := db.Model(&posts).Order("id DESC").Select()
	log.Println(posts, err)
	if err != nil {
		fmt.Println("Error:", err)
	}
	t, _ := template.ParseFiles("templates/layout.html", "templates/index.html")
	data := struct {
		Title string
		Items []Article
	}{
		Title: "Article redactor",
		Items: posts,
	}
	t.Execute(w, data)
	defer db.Close()
}

// InsertPost Create new blog post using form submit
func InsertPost(w http.ResponseWriter, r *http.Request) {
	db := DBConn()
	if r.Method == "POST" {
		text := r.Form.Get("article")
		var preview string
		if len(text) <= 20 {
			preview = text
		} else {
			preview = text[:20] + "..."
		}
		newPost := &Article{
			Preview: preview,
			Article: text,
		}
		err := db.Insert(newPost)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	defer db.Close()
	//http.Redirect(w, r, "/admin", 301)
}
