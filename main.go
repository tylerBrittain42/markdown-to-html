package main

import (
	"github.com/tylerBrittain42/markdown-to-html/parser"
	"os"
	"strings"
)

// stealing from gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := "test/sample.txt"
	name, _ := splitFilename(filename)

	readFile, err := os.Open(filename)
	check(err)
	defer readFile.Close()

	writeFile, err := os.Create(name + ".html")
	check(err)
	defer writeFile.Close()

	parser.Convert(readFile, writeFile)

}



func splitFilename(fileName string) (name string, extension string) {
	parts := strings.Split(fileName, ".")
	return parts[0], parts[1]
}
