package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/tylerBrittain42/markdown-to-html/parser"
)

// stealing from gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var filename string
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print("File name: ")
		fmt.Scan(&filename)
	} else if len(args) == 1 {
		filename = args[0]
	} else {
		//TODO replace with error message
		fmt.Println("too many args")
		os.Exit(1)
	}

	// Creating output folder if it does not exist
	_, err := os.Stat("output")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Creating output folder")
		err = os.Mkdir("output", 0755)
		check(err)
	}

	name, _ := removeExtension(filename)

	readFile, err := os.Open(filename)
	check(err)
	defer readFile.Close()

	writeFile, err := os.Create(filepath.Join("output", filepath.Base(name+".html")))
	check(err)
	defer writeFile.Close()

	parser.Convert(readFile, writeFile)
}

func removeExtension(fileName string) (name string, extension string) {
	parts := strings.Split(fileName, ".")
	return parts[0], parts[1]
}
