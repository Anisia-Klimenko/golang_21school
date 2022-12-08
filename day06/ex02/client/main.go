package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	var body io.Reader
	resp, err := client.Post("https://localhost:8888/admin", "text/plain", body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("response status: %s", resp.Status)
}
