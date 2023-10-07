package main

import (
	"bufio"
	// "fmt"
	"io"
	"os"
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

func convert(readFile io.Reader, writeFile io.Writer) {

	scanner := bufio.NewScanner(readFile)
	writer := bufio.NewWriter(writeFile)
	// blockType := ""

	// foo := ""
	lastType := "first"
	var startBlock, endBlock string

	// newEle is how we determine if we need outer tags in the case of ol and ul
	isOpenList := false

	_, err := writer.WriteString("<html>\n")
	check(err)

	for scanner.Scan() {
		var innerHtml string

		block, text := parser.SplitLine(scanner.Text())
		blockType := parser.GetBlockType(block)
		innerHtml = parser.GetInnerText(text)

		if blockType == "ol" || blockType == "ul" {
			startBlock = parser.OpenTag("li")
			endBlock = parser.CloseTag("li")
		} else if blockType == "" {
			startBlock = parser.OpenTag("p")
			endBlock = parser.CloseTag("p")
		} else {
			startBlock = parser.OpenTag(blockType)
			endBlock = parser.CloseTag(blockType)
		}


		if lastType != blockType && lastType == "ul" || lastType == "ol" {
			_, err = writer.WriteString(parser.CloseTag(lastType) + "\n")
			check(err)
			isOpenList = false
		
		}

		if blockType == "ol" || blockType == "ul" {
				_, err = writer.WriteString(parser.OpenTag(blockType) + "\n")
				check(err)
				isOpenList = true
		}

		htmlLine := createLine(startBlock, innerHtml, endBlock)

		_, err := writer.WriteString(htmlLine)
		check(err)
		_, err = writer.WriteString("\n")
		check(err)
		lastType = blockType

	}
	
	if isOpenList {
		_, err = writer.WriteString(parser.CloseTag(lastType) + "\n")
		check(err)
	}
	_, err = writer.WriteString("</html>\n")
	check(err)
	err = writer.Flush()
	check(err)

}

func createLine(open string, text string, close string) string {
	var htmlBuilder strings.Builder

	htmlBuilder.WriteString(open)
	htmlBuilder.WriteString(text)
	htmlBuilder.WriteString(close)

	return htmlBuilder.String()
}


func parseFilename(name string) (string, string) {
	parts := strings.Split(name, ".")
	return parts[0], parts[1]
}
