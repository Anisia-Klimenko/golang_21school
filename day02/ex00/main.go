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
	inputPath := getPath()
	inputFile, err := os.Open(inputPath)
	if inputPath == "" || err != nil {
		fmt.Println("bad path")
	}
	err = inputFile.Close()
	if err != nil {
		return
	}
	f := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			subPath := strings.TrimPrefix(path, inputPath)
			if subPath != "" && !IsHiddenFile(subPath) {
				if info.Mode()&(1<<2) != 0 {
					if symlinks && info.Mode().Type() == os.ModeSymlink {
						realPath, err := filepath.EvalSymlinks(path)
						if err != nil {
							fmt.Println(path, "-> [broken]")
						} else {
							fmt.Println(path, "->", realPath)
						}
					} else if files && info.Mode().IsRegular() {
						if ext == "" || (ext != "" && filepath.Ext(path) == "."+ext) {
							fmt.Println(path)
						}
					} else if dirs && info.IsDir() {
						fmt.Println(path)
					}
				}
			}
		}
		return nil
	})
	_ = f
}

func IsHiddenFile(filename string) bool {
	if filename[0] == '/' {
		return filename[1] == '.'
	}
	return filename[0] == '.'
}

func getPath() string {
	var path string
	if len(flag.Args()) == 0 {
		path = ""
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
