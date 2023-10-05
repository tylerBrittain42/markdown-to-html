package parser

import (
	"strings"
)

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