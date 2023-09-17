package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type tag struct {
	open    string
	tagType TagType
	close   string
}

type TagType int

const (
	block int = iota
	inline
	list // figure out better name?
)

// stealing from gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// name := "sample.txt"
	// convert(name)

	fmt.Println(strings.Contains("this is a test", "\n"))

	// var foo strings.Builder
	// fmt.Println(foo.String())
	// foo.WriteString("ha")
	// fmt.Println(foo.String())

}

func openTag(tag string) string {
	return "<" + tag + ">"
}

func closeTag(tag string) string {
	return "</" + tag + ">"
}

func parse(line string) string {
	// TODO consider using character array
	var htmlBuilder strings.Builder
	var tokenStack Stack

	// Block Token Check
	first := strings.Split(line, " ")[0]
	block := getBlockType(first)
	if block != "" {
		htmlBuilder.WriteString(openTag(block))
		// plus one  because markdown requires a space
		// is "# " NOT "#"
		line = line[len(first) + 1:]

	}

	// inline check and building
	// Currently just ** and *
	bold := false
	italic := false
	for i := 0; i < len(line); i++ {
		if (i+1) < len(line) && line[i] == '*' && line[i+1] == '*' {
			bold = !bold
			tokenStack.Push("**")
			i++
			if bold {
				htmlBuilder.WriteString(openTag("strong"))
			} else {
				htmlBuilder.WriteString(closeTag("strong"))
			}
		} else if line[i] == '*' {
			italic = !italic
			tokenStack.Push("*")
			if italic {
				htmlBuilder.WriteString(openTag("em"))
			} else {
				htmlBuilder.WriteString(closeTag("em"))
			}
		} else {
			htmlBuilder.WriteByte(line[i])
		}
	}

	if block != "" {
		htmlBuilder.WriteString(closeTag(block))
	}
	// fmt.Println(htmlBuilder.String())
	return htmlBuilder.String()
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

	scanner.Scan()
	// for scanner.Scan() {
	// 	// blockType := getBlockType(scanner.Text())
	// 	contents := getContents(scanner.Text())
	// 	// htmlLine := htmlBuilder(blockType, contents)
	// 	htmlLine := parseLine(contents)
	// 	writer.WriteString(htmlLine)
	// }
	err = writer.Flush()
	check(err)

}

func getBlockType(token string) string {

	blockMarkSymbols := map[string]string{
		"#":      "h1",
		"##":     "h2",
		"###":    "h3",
		"####":   "h4",
		"#####":  "h5",
		"######": "h6",
		"1.":     "ol", //unsure how to handle ordered lists
		"-":      "ul",
		"---":    "br",
		// "\n": "<p>", //unsure how to handle paragraphs
		// block qyotes
		// code
		// DREAM
		// table
		// checklist

	}
	return blockMarkSymbols[token]
}

func parseFilename(name string) (string, string) {
	parts := strings.Split(name, ".")
	return parts[0], parts[1]
}
