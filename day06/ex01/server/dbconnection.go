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
	id   int64  `pg:"articles.id"`
	text string `pg:"articles.article"`
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
	data := context.Get(r, "data")
	t, _ := template.ParseFiles("templates/layout.html", "templates/new.html")
	t.Execute(w, data)
}

// AllPosts Get all blog posts and render template
func AllPosts(w http.ResponseWriter, r *http.Request) {
	db := DBConn()
	var posts []Article
	err := db.Model(&posts).Select()
	log.Println(posts, err)
	if err != nil {
		fmt.Println("Error:", err)
	}
	t, _ := template.ParseFiles("templates/layout.html", "templates/index.html")
	//text := context.Get(r, "article")
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

// InsertPost Create new blog post post using form submit
func InsertPost(w http.ResponseWriter, r *http.Request) {
	db := DBConn()
	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")
		email := context.Get(r, "email").(string)
		post1 := &Article{
			Title:       title,
			Content:     content,
			AuthorEmail: email,
		}
		err := db.Insert(post1)
		if err != nil {
			panic(err)
		}
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
