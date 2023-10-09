package parser

import (
	"bufio"
	"io"
	// "os"
	"strings"
	"bytes"
)

func Convert(readFile io.Reader, writeFileFinal io.Writer) {
	var interimFile bytes.Buffer
	FirstPass(readFile, &interimFile)
	SecondPass(&interimFile, writeFileFinal)
}

// Second pass through to remove empty paragraph tags
func SecondPass(readFile io.Reader, writeFile io.Writer) {

	scanner := bufio.NewScanner(readFile)
	writer := bufio.NewWriter(writeFile)

	for scanner.Scan() {
		if scanner.Text() != "<p></p>" {
			_, err := writer.WriteString(scanner.Text() + "\n")
			check(err)
		}
	}

	err := writer.Flush()
	check(err)


}

func FirstPass(readFile io.Reader, writeFile io.Writer) {

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

		block, text := SplitLine(scanner.Text())
		blockType := GetBlockType(block)
		innerHtml = GetInnerText(text)

		if blockType == "ol" || blockType == "ul" {
			startBlock = OpenTag("li")
			endBlock = CloseTag("li")
		} else if blockType == "" {
			startBlock = OpenTag("p")
			endBlock = CloseTag("p")
		} else {
			startBlock = OpenTag(blockType)
			endBlock = CloseTag(blockType)
		}

		if isOpenList && lastType != blockType && (lastType == "ul" || lastType == "ol") {
			_, err = writer.WriteString(CloseTag(lastType) + "\n")
			check(err)
			isOpenList = false

		} else if !isOpenList && (blockType == "ol" || blockType == "ul") {
			_, err = writer.WriteString(OpenTag(blockType) + "\n")
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
		_, err = writer.WriteString(CloseTag(lastType) + "\n")
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
func SplitLine(line string) (block string, text string) {

	split := strings.SplitN(line, " ", 2)
	if GetBlockType(split[0]) != "" {
		block = split[0]
		text = split[1]
	} else {
		block = ""
		text = line
	}

	return block, text
}

func GetBlockType(token string) string {

	blockMarkSymbols := map[string]string{
		"#":      "h1",
		"##":     "h2",
		"###":    "h3",
		"####":   "h4",
		"#####":  "h5",
		"######": "h6",
		// code
		// DREAM
		// table
		"1.":  "ol", //unsure how to handle ordered lists
		"-":   "ul",
		"---": "br",
		// "\n": "<p>", //unsure how to handle paragraphs
		// block qyotes
		// code
		// DREAM
		// table
		// checklist

	}
	return blockMarkSymbols[token]
}

// Assume block type has been removed
func GetInnerText(line string) string {
	var parsed strings.Builder

	bold := false
	italic := false

	for i := 0; i < len(line); i++ {
		if (i+1) < len(line) && line[i] == '*' && line[i+1] == '*' {
			bold = !bold
			i++
			if bold {
				parsed.WriteString(OpenTag("strong"))
			} else {
				parsed.WriteString(CloseTag("strong"))
			}
		} else if line[i] == '*' {
			italic = !italic
			if italic {
				parsed.WriteString(OpenTag("em"))
			} else {
				parsed.WriteString(CloseTag("em"))
			}
		} else {
			parsed.WriteByte(line[i])
		}
	}
	return parsed.String()
}
func OpenTag(tagVal string) string {
	tag := ""
	if tagVal == "li" {
		tag += "\t"
	}
	tag += "<" + tagVal + ">"
	return tag
}

func CloseTag(tag string) string {
	return "</" + tag + ">"
}

// stealing from gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}
