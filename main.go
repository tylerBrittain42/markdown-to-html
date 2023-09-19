package main

import (
	"bufio"
	"io"
	// "fmt"
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
	name, _ := parseFilename(filename)

	readFile, err := os.Open(filename)
	check(err)
	defer readFile.Close()

	writeFile, err := os.Create(name + ".html")
	check(err)
	defer writeFile.Close()

	convert(readFile, writeFile)

}

func openTag(tag string) string {
	return "<" + tag + ">"
}

func closeTag(tag string) string {
	return "</" + tag + ">"
}

func buildLine(line string) string {
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
		line = line[len(first)+1:]

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
	return htmlBuilder.String()
}

// SPLIT THIS INTO MULTIPLE FILES FOR TESTING
// SEE STACK OVERFLOW PAGE
func convert(readFile io.Reader, writeFile io.Writer) {

	scanner := bufio.NewScanner(readFile)
	writer := bufio.NewWriter(writeFile)

	for scanner.Scan() {
		htmlLine := buildLine(scanner.Text())
		_, err := writer.WriteString(htmlLine + "\n")
		check(err)
	}

	err := writer.Flush()
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
