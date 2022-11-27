package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

//func main() {
//	cmd := exec.Command("wc", "-l")
//	//stdout, err := cmd.Output()
//	//
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//	return
//	//}
//	//
//	//// Print the output
//	//fmt.Println(string(stdout))
//
//	//cmd := exec.Command("cat")
//	stdin, err := cmd.StdinPipe()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	go func() {
//		defer stdin.Close()
//		io.WriteString(stdin, "an old falcon")
//	}()
//
//	out, err := cmd.CombinedOutput()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("%s\n", out)
//}

var _ = io.WriteString

func main() {
	cmd := exec.Command("wc", "-m")

	/*
	   stdin, err := cmd.StdinPipe()
	   if err != nil {
	     log.Fatal(err)
	   }
	   go func() {
	     defer stdin.Close()
	     io.WriteString(stdin, "an old falcon\n")
	   }()
	*/
	cmd.Stdin = os.Stdin

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
