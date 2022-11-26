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
	if path.Ext(*fileName) == ".xml" {
		err, content := XMLname(*fileName).read()
		if err != nil {
			fmt.Println(err)
			return
		}
		//XMLname(*fileName).print(content)
		fmt.Println(XMLname(*fileName).convert(content))
	} else if path.Ext(*fileName) == ".json" {
		err, content := JSONname(*fileName).read()
		if err != nil {
			fmt.Println(err)
			return
		}
		//JSONname(*fileName).print(content)
		fmt.Println(JSONname(*fileName).convert(content))
	} else {
		fmt.Println("usage: ./readDB -f <filename>.xml")
		fmt.Println("       ./readDB -f <filename>.json")
	}
}
