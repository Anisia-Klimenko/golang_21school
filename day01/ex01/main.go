package main

import (
	"flag"
	"fmt"
	"path"
)

const (
	BLUE   = "\033[1;34m"
	GRN    = "\033[1;32m"
	RED    = "\033[0;31m"
	VIOLET = "\033[0;35m"
	YELLOW = "\033[1;33m"
	TICK   = "\xE2\x9C\x94"
	END    = "\033[0m"
)

func main() {
	oldDB := flag.String("old", "", "old database")
	newDB := flag.String("new", "", "new database")
	flag.Parse()
	var dbreaderOld DBReader
	var dbreaderNew DBReader

	if path.Ext(*oldDB) == ".xml" && path.Ext(*newDB) == ".json" {
		dbreaderOld = XMLname(*oldDB)
		err, old := dbreaderOld.read()
		if err != nil {
			fmt.Println("old database is broken")
			return
		}
		dbreaderNew = JSONname(*newDB)
		err, new := dbreaderNew.read()
		if err != nil {
			fmt.Println("new database is broken")
			return
		}
		compare(old, new)
	} else if path.Ext(*newDB) == ".xml" && path.Ext(*oldDB) == ".json" {
		dbreaderOld = JSONname(*oldDB)
		err, old := dbreaderOld.read()
		if err != nil {
			fmt.Println("old database is broken")
			return
		}
		dbreaderNew = XMLname(*newDB)
		err, new := dbreaderNew.read()
		if err != nil {
			fmt.Println("new database is broken")
			return
		}
		compare(old, new)
	} else {
		fmt.Println("wrong extension")
	}
}
