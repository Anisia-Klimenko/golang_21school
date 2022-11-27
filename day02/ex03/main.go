package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	newDir := flag.String("a", "", "output path of archive files")

	_ = newDir
	for count, file := range os.Args {
		if count != 0 {
			fmt.Println(file)
			wg.Add(1)
			file := file
			go func() {
				var st syscall.Stat_t
				if err := syscall.Stat(file, &st); err != nil {
					log.Fatal(err)
				}
				fmt.Println(st.Mtimespec.Sec, strings.TrimSuffix(file, filepath.Ext(file))+"_"+strconv.FormatInt(st.Mtimespec.Sec, 10))
				wg.Done()
			}()
		}
	}
	wg.Wait()
}
