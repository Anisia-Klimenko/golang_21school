package main

import (
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

var rl = rate.NewLimiter(rate.Every(time.Second), 100)

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rl.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Printf("Server started")
	initCredentials()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}
