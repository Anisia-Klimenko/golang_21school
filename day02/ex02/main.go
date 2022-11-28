package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

//stdout, err := cmd.Output()
//
//if err != nil {
//	fmt.Println(err.Error())
//	return
//}
//
//// Print the output
//fmt.Println(string(stdout))

//cmd := exec.Command("cat")

var _ = io.WriteString

func main() {
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	var input []byte
	_, _ = io.PipeReader.Read(io.PipeReader{}, input)

	//for _, line := range strings.Split(strings.TrimSuffix(string(input), "\n"), "\n") {
	//	cmd.Stdin = strings.NewReader(line)
	//fmt.Println(input)
	cmd.Stdin = strings.NewReader(string(input))
	//cmd.Stdin = os.Stdin

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
	//}

	//cmd.Stdin = os.Stdin
	//
	//out, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%s\n", out)
}

//echo -e "main.go\ngo.mod" | xargs cat
