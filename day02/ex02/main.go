package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[2:]
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("error: io.ReadAll")
		return
	}
	for _, line := range strings.Split(strings.TrimSuffix(string(input), "\n"), "\n") {
		cmd := exec.Command(os.Args[1], append(args, line)...)
		cmd.Stdin = strings.NewReader(line)

		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", out)
	}
}

//echo -e "main.go\ngo.mod" | xargs cat
//echo -e "main.go\ngo.mod" | xargs wc -l
// echo -e "/Users/acristin/golang_21school/day02/\n." | xargs ls -l
