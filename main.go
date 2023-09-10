package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type tag struct {
	open  string
	close string
}

// stealing from gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// REFERENCE: this is how to make a map
	// markToHtml := make(map[string]tag)
	// markToHtml := map[string]tag{
	// 	"#":      {"<h1>", "</h1>"},
	// 	"##":     {"<h2>", "</h2>"},
	// 	"###":    {"<h3>", "</h3>"},
	// 	"####":   {"<h4>", "</h4>"},
	// 	"#####":  {"<h5>", "</h5>"},
	// 	"######": {"<h6>", "</h6>"}}
	// simultaneousHandlerExample("read_input.md")
	makeCapsLock("read_input.md")
}

func makeCapsLock(fileName string) {
	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	writeFile, err := os.Create("new_" + fileName)
	check(err)
	defer writeFile.Close()

	scanner := bufio.NewScanner(readFile)
	writer := bufio.NewWriter(writeFile)

	for scanner.Scan() {
		_,err = writer.WriteString(strings.ToUpper(scanner.Text() + "\n"))
		check(err)
	}
	err = writer.Flush()
	check(err)
}

// this actually has a cool example of how the file handlers interact
// if you scan after the flush then you will recieve the updated text
// however if you scan beforehand, that scanned original text will remain
func simultaneousHandlerExample(fileName string) {
	// setup of files
	filePath := fileName
	fileWrite, err := os.OpenFile(filePath, os.O_RDWR, 0777)
	check(err)
	defer fileWrite.Close()
	fileRead, err := os.Open(filePath)
	check(err)
	defer fileRead.Close()

	scanner := bufio.NewScanner(fileRead)
	scanner.Scan()

	writer := bufio.NewWriter(fileWrite)
	_, err = writer.Write([]byte("This statement is from the file handler\n"))
	check(err)
	// fmt.Println(writer)
	err = writer.Flush()
	check(err)
	// scanner.Scan()
	fmt.Println(scanner.Text())
}
