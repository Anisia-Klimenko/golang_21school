package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/gorilla/context"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

// ShowPost Update post details
func ShowPost(w http.ResponseWriter, r *http.Request) {
	db := DBConn()
	nId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	// Select user by primary key.
	post := &Article{Id: nId}
	err = db.Select(post)
	//log.Printf("showPost -> %s", post.Preview)
	if err != nil {
		fmt.Println("Error:", err)
	}
	var tplMap = map[string]interface{}{
		"Article": template.HTML(post.Article),
		"Preview": post.Preview,
		"Id":      post.Id,
	}
	t, _ := template.ParseFiles("templates/layout.html", "templates/show.html")
	t.Execute(w, tplMap)
	defer db.Close()
}

// NewPost Return new blog Post html form on GET
func NewPost(w http.ResponseWriter, r *http.Request) {
	data := context.Get(r, "article")
	log.Println("POST /admin/new NewPost")
	t, _ := template.ParseFiles("templates/layout.html", "templates/new.html")
	t.Execute(w, data)
	http.Redirect(w, r, "/admin", 301)
}

func NewGet(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /admin/new NewPost")
	t, _ := template.ParseFiles("templates/layout.html", "templates/new.html")
	t.Execute(w, nil)
}

func SendLogo(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadFile("resources/amazing_logo.png")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(buf)
}

// AllPosts Get all blog posts and render template
func AllPosts(w http.ResponseWriter, r *http.Request) {
	db := DBConn()
	var posts []Article
	err := db.Model(&posts).Order("id DESC").Select()
	log.Println("GET /admin AllPosts")
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

var (
	boldItalicReg = regexp.MustCompile(`\*\*\*(.*?)\*\*\*`)
	boldReg       = regexp.MustCompile(`\*\*(.*?)\*\*`)
	italicReg     = regexp.MustCompile(`\*(.*?)\*`)
	strikeReg     = regexp.MustCompile(`\~\~(.*?)\~\~`)
	underscoreReg = regexp.MustCompile(`__(.*?)__`)
	//anchorReg     = regexp.MustCompile(`\[(.*?)\]\((.*?)\)[^\)]`)
	anchorReg = regexp.MustCompile(`(https://)(?:\S+)(/*)`)
	//(http.://([^\s]+))
	escapeReg     = regexp.MustCompile(`^\>(\s|)`)
	blockquoteReg = regexp.MustCompile(`\&gt\;(.*?)$`)
	backtipReg    = regexp.MustCompile("`(.*?)`")

	h1Reg = regexp.MustCompile(`^#(\s|)(.*?)$`)
	h2Reg = regexp.MustCompile(`^##(\s|)(.*?)$`)
	h3Reg = regexp.MustCompile(`^###(\s|)(.*?)$`)
	h4Reg = regexp.MustCompile(`^####(\s|)(.*?)$`)
	h5Reg = regexp.MustCompile(`^#####(\s|)(.*?)$`)
	h6Reg = regexp.MustCompile(`^######(\s|)(.*?)$`)
)

func NewMarkdown(input io.Reader) string {

	buf := bytes.NewBuffer(nil)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			buf.WriteByte('\n')
			continue
		}

		// wrap bold and italic text in "<b>" and "<i>" elements
		line = boldItalicReg.ReplaceAll(line, []byte(`<b><i>$1</i></b>`))
		line = boldReg.ReplaceAll(line, []byte(`<b>$1</b>`))
		line = italicReg.ReplaceAll(line, []byte(`<i>$1</i>`))
		// wrap strikethrough text in "<s>" tags
		line = strikeReg.ReplaceAll(line, []byte(`<s>$1</s>`))
		// wrap underscored text in "<u>" tags
		line = underscoreReg.ReplaceAll(line, []byte(`<u>$1</u>`))
		// convert links to anchor tags
		line = anchorReg.ReplaceAll(line, []byte(`<a href=\"$0\">$0</a>`))
		// escape and wrap blockquotes in "<blockquote>" tags
		line = escapeReg.ReplaceAll(line, []byte(`&gt;`))
		line = blockquoteReg.ReplaceAll(line, []byte(`<blockquote>$1</blockquote>`))
		// wrap the content of backticks inside "<code>" tags
		line = backtipReg.ReplaceAll(line, []byte(`<code>$1</code>`))
		// convert headings
		if line[0] == '#' {

			count := bytes.Count(line, []byte(`#`))
			switch count {
			case 1:
				line = h1Reg.ReplaceAll(line, []byte(`<h1>$2</h1>`))
			case 2:
				line = h2Reg.ReplaceAll(line, []byte(`<h2>$2</h2>`))
			case 3:
				line = h3Reg.ReplaceAll(line, []byte(`<h3>$2</h3>`))
			case 4:
				line = h4Reg.ReplaceAll(line, []byte(`<h4>$2</h4>`))
			case 5:
				line = h5Reg.ReplaceAll(line, []byte(`<h5>$2</h5>`))
			case 6:
				line = h6Reg.ReplaceAll(line, []byte(`<h6>$2</h6>`))
			}
		}
		buf.Write(line)
		buf.WriteByte('\n')
	}
	return buf.String()
}

// InsertPost Create new blog post using form submit
func InsertPost(w http.ResponseWriter, r *http.Request) {
	db := DBConn()
	if r.Method == "POST" {
		text := r.Form.Get("article")
		if len(text) != 0 {
			var preview string
			if len(text) <= 20 {
				preview = text
			} else {
				preview = text[:20] + "..."
			}
			text = NewMarkdown(strings.NewReader(text))
			//log.Println(text)
			newPost := &Article{
				Preview: preview,
				Article: text,
			}
			err := db.Insert(newPost)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
	defer db.Close()
	//http.Redirect(w, r, "/admin", 301)
}
