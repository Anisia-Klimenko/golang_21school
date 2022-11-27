package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var symlinks = false
var dirs = false
var files = false
var ext = ""

func main() {
	getFlags()
	//files, err := ioutil.ReadDir(getPath())
	//if err != nil {
	//	log.Fatal(err)
	//}

	f := filepath.Walk(getPath(), func(path string, info os.FileInfo, err error) error {
		if err == nil {
			subPath := strings.TrimPrefix(path, getPath())
			if subPath != "" && !IsHiddenFile(subPath) {

				if dirs && info.IsDir() {
					fmt.Println(path)
				}

			}
		}
		return nil
	})

	_ = f
	//_ = files

	//for _, file := range files {
	//	fmt.Println(file.Name(), file.IsDir())
	//}
	fmt.Println(flag.Args())
}

func IsHiddenFile(filename string) bool {
	return filename[0] == '.'
}

func getPath() string {
	var path string
	if len(flag.Args()) == 0 {
		path = "./"
	} else {
		path = flag.Args()[0]
	}
	return path
}

func getFlags() {
	flag.BoolVar(&symlinks, "sl", true, "symlinks")
	flag.BoolVar(&dirs, "d", true, "directories")
	flag.BoolVar(&files, "f", true, "files")
	flag.StringVar(&ext, "ext", "", "extensions")
	flag.Parse()
	if !isFlagPassed("sl") {
		symlinks = false
	}
	if !isFlagPassed("d") {
		dirs = false
	}
	if !isFlagPassed("f") {
		files = false
	}
	if !isFlagPassed("ext") || !isFlagPassed("f") {
		ext = ""
	}
	if !symlinks && !dirs && !files {
		symlinks = true
		dirs = true
		files = true
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
