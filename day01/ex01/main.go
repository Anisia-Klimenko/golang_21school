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

	if path.Ext(*oldDB) == ".xml" && path.Ext(*newDB) == ".json" {
		err, old := XMLname(*oldDB).read()
		if err != nil {
			fmt.Println("old database is broken")
			return
		}
		err, new := JSONname(*newDB).read()
		if err != nil {
			fmt.Println("new database is broken")
			return
		}
		//JSONname(*newDB).print(new)
		//fmt.Println("=========================")
		//XMLname(*oldDB).print(old)
		//_ = old
		compare(old, new)
	} else if path.Ext(*newDB) == ".xml" && path.Ext(*oldDB) == ".json" {
		err, old := JSONname(*oldDB).read()
		if err != nil {
			fmt.Println("old database is broken")
			return
		}
		err, new := XMLname(*newDB).read()
		if err != nil {
			fmt.Println("new database is broken")
			return
		}
		compare(old, new)
	} else {
		fmt.Println("wrong extension")
	}
}
