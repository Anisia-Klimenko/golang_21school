package main

import (
	"archive/tar"
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
	flag.Parse()

	_ = newDir
	fmt.Println("a", *newDir)
	if *newDir != "" {
		fi, err := os.Stat(*newDir)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !fi.Mode().IsDir() {
			fmt.Println("-a: no such directory")
			return
		}
		if (*newDir)[len(*newDir)-1] != '/' {
			*newDir += "/"
		}
	}
	for _, file := range flag.Args() {
		fmt.Println(file)
		wg.Add(1)
		file := file
		go func() {
			var st syscall.Stat_t
			if err := syscall.Stat(file, &st); err != nil {
				log.Fatal(err)
			}
			filename := strings.TrimSuffix(file, filepath.Ext(file))
			tarPath := *newDir + filename + "_" + strconv.FormatInt(st.Mtimespec.Sec, 10)
			tarFile, err := os.Create(tarPath + ".tar.gz")
			if err != nil {
				log.Fatal(err)
			}
			defer tarFile.Close()
			tw := tar.NewWriter(tarFile)
			defer tw.Close()
			content, _ := os.ReadFile(file)
			hdr := &tar.Header{
				Name: filename,
				Mode: 0600,
				Size: int64(len(content)),
			}
			if err := tw.WriteHeader(hdr); err != nil {
				panic(err)
			}
			if _, err := tw.Write(content); err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
