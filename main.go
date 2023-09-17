package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type tag struct {
	open string
	tagType TagType
	close string
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
	// var s Stack
	// fmt.Println(s)
	// s.Push("1asdf")
	// s.Push("2fast")
	// s.Push("3asdf")
	// fmt.Println(s)
	// fmt.Println("peaked: " + s.Peak())
	// s.Pop()
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())

}

func parse(line string) Stack {
	blockSymbols := map[string]string{
		"#":      "h1",
		"##":     "h2",
		"###":    "h3",
		"####":   "h4",
		"#####":  "h5",
		"######": "h6",
		"1.":     "ol", //unsure how to handle ordered lists
		"-":      "ul",
		"---":    "br",
		// "**": "bold",
		// "*": "italic",
		// "\n": "<p>", //unsure how to handle paragraphs
		// block qyotes
		// code
		// DREAM
		// table
		// checklist

	} 
	// LETS SPLIT BLOCK PARSING FROM inline
	// 1. BLOCK PARSE FIRST TOKEN
	// 2. CONTINUE WITH PARSING EVERYTING ELSE
	// WHY THE FUCK AM I WORRYING ABOUT UL AND OL
	// IT WILL ONLY MATTER ON THE REVERSE
	// LMAOOOOOOOOOOOOOOO
	var tokenStack Stack

	// for _,v := range parts {
	// 	if MarkSymbols[v] != "" {
	// 		tokenStack.Push(v)
	// 	} else if v[0] == '*' && v[1] == '*' {
	// 		tokenStack.Push("**")
	// 	} else if v[0] == '*' && v[1] != '*'{
	// 		tokenStack.Push("*")
	// 	}
	// 		
	// }

	// Block Token Check
	blockToken := strings.Split(line, " ")[0]
	fmt.Println(blockToken)
	if blockSymbols[blockToken] != "" {
		tokenStack.Push(blockToken)
	}

	// inline check
	// Currently just ** and *
	for i := 0; i < len(line); i++{
		if (i+1) < len(line) && line[i] == '*' && line[i+1] == '*' {
			tokenStack.Push("**")
			i++
		} else if line[i] == '*'{
			tokenStack.Push("*")
		}
	}
			
	return tokenStack
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
		// blockType := getBlockType(scanner.Text())
		contents := getContents(scanner.Text())
		// htmlLine := htmlBuilder(blockType, contents)
		htmlLine := parseLine(contents)
		writer.WriteString(htmlLine)
	}
	err = writer.Flush()
	check(err)

}

func parseLine(line string) string {
	// tokenStack = []string
	// tokenStack =
	// var tokenStack Stack = []{"this","is"}


	return "a"
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
		"---":    "br",
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
