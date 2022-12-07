package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	// Replace string with the specified regexp
	// Using ReplaceAllString() method
	m1 := regexp.MustCompile(`x(?:\S+)y`)

	fmt.Println(m1.ReplaceAllString("xy--xpppyxxppxy-", "B"))
	fmt.Println(m1.ReplaceAllString("xy--xpppyxxppxy--", "$1"))
	fmt.Println(m1.ReplaceAllString("xy--xpppyxxppxy-", "$1P"))
	fmt.Println(m1.ReplaceAllString("xy--xpppyxxppxy-", "${1}Q"))

	fmt.Println(m1.ReplaceAllString("xppy--xpppyxxppxy--", "$1"))

	anchorReg := regexp.MustCompile(`(https://)(?:\S+)(/*)`)
	//anchorReg := regexp.MustCompile(`\[(.*?)\]\((.*?)\)[^\)]`)
	line := "google.com\nhttps://google.com\nhttps://google.com\ngoogle.com/search\nhttps://harm-smits.github.io/42docs/\nhttps://harm-smits.github.io/42docs\nhttps://harm-smits.github.io/42docs/123/123\ngoogle.com.id\n\nhttps://example.de \nhttps://example.de/\nhttps://example.de/home\nhttps://example.de/home/\nhttps://example.de/home some text that should not be extracted\nhttps://abc.example.de\nhttps://abc.example.de/\nhttps://abc.example.de/home\nhttps://abc.example.de/home\nhttps://abc.example.de/home some text that should not be extracted"
	//line := "2D игра на С, с использованием библиотеки MinilibX (подробнее про MinilibX: https://harm-smits.github.io/42docs/)\n"
	for _, l := range strings.Split(line, "\n") {

		//fmt.Println(string(anchorReg.ReplaceAll([]byte(l), []byte("<a href=\""+l+"\">"+l+"</a>"))))
		fmt.Println(string(anchorReg.ReplaceAll([]byte(l), []byte(`<a href=\"$0\">$0</a>`))))
	}

}
