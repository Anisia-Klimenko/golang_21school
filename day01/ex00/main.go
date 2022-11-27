package main

import (
	"flag"
	"fmt"
	"path"
)

func main() {
	var fileName *string
	fileName = flag.String("f", "", "input file name")
	flag.Parse()
	var dbreader DBReader

	if path.Ext(*fileName) == ".xml" {
		dbreader = XMLname(*fileName)
		err, content := dbreader.read()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(dbreader.convert(content))
	} else if path.Ext(*fileName) == ".json" {
		dbreader = JSONname(*fileName)
		err, content := dbreader.read()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(dbreader.convert(content))
	} else {
		fmt.Println("usage: ./readDB -f <filename>.xml")
		fmt.Println("       ./readDB -f <filename>.json")
	}
}
