package main

import (
	"bufio"
	"fmt"
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
	var newEle bool

	_, err := writer.WriteString("<html>\n")
	check(err)

	for scanner.Scan() {
		//TODO: remove me
		fmt.Println(scanner.Text())

		var innerHtml string

		block, text := parser.SplitLine(scanner.Text())
		blockType := parser.GetBlockType(block)

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

		innerHtml = parser.GetInnerText(text)

		// hangle open tag
		// handle close tag
		// build html
		// output

		// innerText := strings.Split(scanner.Text(),blockType)

		if lastType == "first" {
			newEle = true

			// }else if lastType != blockType && lastType != "first" {
		} else if lastType != blockType && lastType == "ul" || lastType == "ol" {
			_, err = writer.WriteString(parser.CloseTag(lastType) + "\n")
			check(err)
			newEle = true
		}

		if newEle {
			switch blockType {
			case "ul":
				// fmt.Println("in switch")
				_, err = writer.WriteString(parser.OpenTag(blockType) + "\n")
				check(err)
			case "ol":
				// fmt.Println("in switch")
				_, err = writer.WriteString(parser.OpenTag(blockType) + "\n")
				check(err)
			// headers
			default:
				// fmt.Println("in switch")
			}

			check(err)
		}

		// // first line
		// if lastType == "" {
		// 	// Normal type
		// 	if blockType != "" && blockType != "ol" && blockType != "ul" {
		// 		startBlock = parser.OpenTag(blockType)
		// 		endBlock = parser.CloseTag(blockType)
		// 	} else if blockType == "ol" {
		// 		startBlock = parser.OpenTag(blockType)
		// 		endBlock = parser.CloseTag(blockType)
		//
		// }
		htmlLine := createLine(startBlock, innerHtml, endBlock)

		// fmt.Println("startBlock:." + startBlock + ".")
		// fmt.Println("innerHtml:." + innerHtml + ".")
		// fmt.Println("endBlock:." + endBlock + ".")
		// fmt.Println("htmlLine:." + htmlLine + ".")

		_, err := writer.WriteString(htmlLine)
		check(err)
		_, err = writer.WriteString("\n")
		check(err)
		lastType = blockType

		newEle = false
	}

	_, err = writer.WriteString("</html>\n")
	check(err)
	err = writer.Flush()
	check(err)

}

// if lastType != blockType{
// 	writer.WriteString(parser.CloseTag(lastType) + "\n")
// 	lastType = "none"
// }
//
// if blockType == "" {
// 	lastType = "p"
// } else if blockType == "li" || blockType == "ul"{
// 	lastType = blockType
// }
// if blockType != "" && blockType != "li" && blockType != "ul" {
// startBlock = parser.OpenTag(blockType)
// endBlock = parser.CloseTag(blockType)
// }

func createLine(open string, text string, close string) string {
	var htmlBuilder strings.Builder

	htmlBuilder.WriteString(open)
	htmlBuilder.WriteString(text)
	htmlBuilder.WriteString(close)

	return htmlBuilder.String()
}

func getStartBlockType(block string) string {
	startBlock := ""

	// switch block {
	// case "li":

	return startBlock
}

//	func parser.getBlockType(foo string) string {
//		return ""
//	}
func handleEndBlockType(foo string) string {
	return ""
}

// NEED
// BASIC PARAGRAPH HANDLING (BLOCK TO NON BLOCK TO BLOCK)
// PARAGRAPH TESTING
// func convert(readFile io.Reader, writeFile io.Writer) {
//
// 	scanner := bufio.NewScanner(readFile)
// 	writer := bufio.NewWriter(writeFile)
// 	var isPara bool
// 	wasClosed := true
// 	isFirst := true
//
// 	for scanner.Scan() {
// 		potentBlock := parser.getBlockType(strings.Split(scanner.Text(), " ")[0])
// 		if isFirst  && potentBlock == ""{
// 			_, err := writer.WriteString(parser.OpenTag("p"))
// 			check(err)
// 			wasClosed = true
// 		} else if isFirst && potentBlock != "" {
// 		} else if potentBlock == "" {
// 			if !isPara {
// 				_, err := writer.WriteString("\n" + parser.OpenTag("p"))
// 				check(err)
// 				isPara = true
// 				wasClosed = false
// 			}
// 		} else if potentBlock != "" && isPara {
// 			_, err := writer.WriteString(parser.CloseTag("p") + "\n")
// 			check(err)
// 			wasClosed = true
// 		} else {
// 			_, err := writer.WriteString("\n")
// 			check(err)
// 		}
// 		isFirst = false
//
// 		htmlLine := buildLine(scanner.Text())
// 		_, err := writer.WriteString(htmlLine)
// 		check(err)
// 		fmt.Println("HTML:." + htmlLine + ".")
// 	}
// 	if !wasClosed {
// 		_, err := writer.WriteString(parser.CloseTag("p") + "\n")
// 		check(err)
// 	} else {
// 		_, err := writer.WriteString("\n")
// 		check(err)
// 	}
// 	err := writer.Flush()
// 	check(err)
//
// }

func parseFilename(name string) (string, string) {
	parts := strings.Split(name, ".")
	return parts[0], parts[1]
}
