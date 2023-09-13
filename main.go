package main

import (
	"bufio"
	"os"
	"strings"
)

type tag struct {
	open string
	// close string
}

// stealing from gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	name := "sample.txt"
	convert(name)

}

// no parsing for now
func convert(filename string) {
	name, extension := parseFilename(filename)

	readFile, err := os.Open(name + "." + extension)
	check(err)
	defer readFile.Close()

	writeFile, err := os.Create(name + "_markdown." + extension)
	check(err)
	defer writeFile.Close()

	scanner := bufio.NewScanner(readFile)
	writer := bufio.NewWriter(writeFile)

	for scanner.Scan() {
		blockType := getBlockType(scanner.Text())
		contents := getContents(scanner.Text())
		htmlLine := htmlBuilder(blockType, contents)
		writer.WriteString(htmlLine)
	}
	err = writer.Flush()
	check(err)

}
func htmlBuilder(blockType string, contents string) string {
	blockStart := "<" + blockType + ">"
	blockEnd := "</" + blockType + ">"
	html := blockStart + contents + blockEnd
	return html
}
func getContents(line string) string {
	return strings.SplitN(line, " ", 2)[1]
}
func getBlockType(line string) string {

	blockMarkSymbols := map[string]string{
		"#":      "h1",
		"##":     "h2",
		"###":    "h3",
		"####":   "h4",
		"#####":  "h5",
		"######": "h6",
		"1.":     "ol", //unsure how to handle ordered lists
		"-":      "ul",
		"---":    "/br",
		// "\n": "<p>", //unsure how to handle paragraphs
		// block qyotes
		// code
		// DREAM
		// table
		// checklist

	}
	token := strings.Split(line, " ")[0]
	blockType := blockMarkSymbols[token]

	return blockType
}

func parseFilename(name string) (string, string) {
	parts := strings.Split(name, ".")
	return parts[0], parts[1]
}
