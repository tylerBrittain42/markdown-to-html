package main

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
	"github.com/tylerBrittain42/markdown-to-html/parser"
)

func TestGetBlockType(t *testing.T) {
	tests := []struct {
		input        string
		expectedOpen string
	}{
		{"# this is a heading", "h1"},
		{"## this is a heading", "h2"},
		{"### this is a heading", "h3"},
		{"#### this is a heading", "h4"},
		{"##### this is a heading", "h5"},
		{"###### this is a heading", "h6"},
		{"1. this is an ordered list", "ol"},
		{"- this is an unordered list", "ul"},
		{"--- this is a line break", "br"},
	}
	fmt.Println("\nTest TestGetBlockType")
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.input)
		fmt.Println(testname)
	}
}

func TestSplitLine(t *testing.T) {
	tests := []struct {
		inputString   string
		expectedBlock string
		expectedBody  string
	}{
		{"# This has a header", "#", "This has a header"},
		{"- This has a bullet", "-", "This has a bullet"},
		{"This has a paragraph", "", "This has a paragraph"},
	}
	fmt.Println("\nTest TestSplitLine")
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			actualBlock, actualBody := parser.SplitLine(tt.inputString)
			if actualBlock != tt.expectedBlock || actualBody != tt.expectedBody {
				t.Errorf("got .%s. .%s., wanted .%s. .%s.", actualBlock, actualBody, tt.expectedBlock, tt.expectedBody)
			}
		})
	}
}

func TestInnerText(t *testing.T) {
	tests := []struct {
		inputString    string
		expectedString string
	}{
		{"This is a single line", "This is a single line"},
		{"This has **underlined** characters", "This has <strong>underlined</strong> characters"},
		{"This has *italic* characters", "This has <em>italic</em> characters"},
		{"This has **underlined** characters and *italic ones*", "This has <strong>underlined</strong> characters and <em>italic ones</em>"},
	}
	fmt.Println("\nTest TestInnerText")

	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := parser.GetInnerText(tt.inputString)
			if ans != tt.expectedString {
				t.Errorf("got .%s., wanted .%s.", ans, tt.expectedString)
			}
		})
	}
}

func TestOpenTag(t *testing.T) {
	tests := []struct {
		inputString    string
		expectedString string
	}{
		{"a", "<a>"},
	}
	fmt.Println("\nTest TestOpenTag")
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := parser.OpenTag(tt.inputString)
			if ans != tt.expectedString {
				t.Errorf("got .%s., wanted .%s.", ans, tt.expectedString)
			}
		})
	}
}

func TestCloseTag(t *testing.T) {
	tests := []struct {
		inputString    string
		expectedString string
	}{
		{"a", "</a>"},
	}
	fmt.Println("\nTest TestOpenTag")
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := parser.CloseTag(tt.inputString)
			if ans != tt.expectedString {
				t.Errorf("got .%s., wanted .%s.", ans, tt.expectedString)
			}
		})
	}
}

type convertTest struct {
	caseName   string
	inputFile  []string
	outputFile []string
}

func createTestCases() []convertTest{
	caseName := []string{
		"1. Single block line",
		"2. Single paragraph line",
		"3. Single ordered list line",
		"4. Multiple block line",
		"5. Multiple block line with italics/bold",
		"6. Multiple block and paragraph line",
		"7. Multiple block and ordered list",
		"8. Multiple block, paragraph, and ordered list with italics and bold",
	}
	inputCases := [][]string{
		// 1
		{
			"# This is a single block line",
		},
		// 2
		{
			"This is a single paragraph line",
		},
		// 3
		{
			"1. This is the first element of a list",
		},
		// 4
		{
			"# This is the first line",
			"## This is the second line",
		},
		// 5
		{
			"# This is the *first* line",
			"## This is the **second** line",
		},
		// 6
		{
			"# This is the first line",
			"This is part of a paragraph",
		},
		// 7
		{
			"# This is the first line",
			"- This is part of a list",
		},
		// 8
		{
			"# This is the *first* line",
			"- This is part of a **list**",
			"This is part of a paragraph",
		},
	}
	expectedOutputs := [][]string{
		{
			// 1
			"<html>",
			"<h1>This is a single block line</h1>",
			"</html>",
		}, {
		// 2
			"<html>",
			"<p>This is a single paragraph line</p>",
			"</html>",
		},{
		// 3
			"<html>",
			"<ol>",
			"\t<li>This is a single line</li>",
			"</ol>",
			"</html>",
		},{
		// 4
			"<html>",
			"<h1>This is the first line</h1>",
			"<h2>This is the second line</h2>",
			"</html>",
		},{
		// 5
			"<html>",
			"<h1>This is the <strong>first</strong> line</h1>",
			"<h2>This is the <em>second</em> line</h2>",
			"</html>",
		},{
		// 6
			"<html>",
			"<h1>This is the first line</h1>",
			"<p> This is part of a paragraph</p>",
			"</html>",
		},{
		// 7
			"<html>",
			"<h1>This is the first line</h1>",
			"<ul>",
			"\t<li>This is part of a list</li>",
			"</ul>",
			"</html>",
		},{
		// 8
			"<html>",
			"<h1>This is the <strong>first</strong> line</h1>",
			"<ul>",
			"\t<li>This is part of a <em>list</em></li>",
			"</ul>",
			"<p>This is part of a paragraph</p>",
			"</html>",
		},
	}
	tests := []convertTest{}
	for i := range caseName {
		tests = append(tests, convertTest{caseName: caseName[i], inputFile: inputCases[i], outputFile: expectedOutputs[i]})
	}
	return tests
}
func TestConvert(t *testing.T) {

	// Input and output are separate variables because the initial <html> tag will make the lines off

	
	tests := createTestCases()

	fmt.Println("\nTest: TestConvert")
	for _, tt := range tests {
		// load buffer
		var readBuffer, writerBuffer bytes.Buffer
		for _, line := range tt.inputFile {
			readBuffer.WriteString(string(line) + "\n")
		}

		testname := fmt.Sprintf("Case: %s", tt.caseName)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			convert(&readBuffer, &writerBuffer)
			checkOutput := bufio.NewScanner(&writerBuffer)
			for j, mock := range tt.outputFile {
				checkOutput.Scan()
				if string(mock) != checkOutput.Text() {
					t.Errorf("%v) got .%s., wanted .%s.", j, checkOutput.Text(), string(mock))
					break
				}
			}
		})
	}

}

// func TestBuildLine(t *testing.T) {
// 	tests := []struct {
// 		inputString    string
// 		expectedString string
// 	}{
// 		{"# This is a single line", "<h1>This is a single line</h1>"},
// 		{"## This is a single line", "<h2>This is a single line</h2>"},
// 		{"### This has **underlined** characters", "<h3>This has <strong>underlined</strong> characters</h3>"},
// 		{"## This has **underlined** characters and *italic ones*", "<h2>This has <strong>underlined</strong> characters and <em>italic ones</em></h2>"},
// 	}
// 	for _, tt := range tests {
// 		testname := fmt.Sprintf("input: %s", tt.inputString)
// 		fmt.Println(testname)
// 		t.Run(testname, func(t *testing.T) {
// 			ans := buildLine(tt.inputString)
// 			if ans != tt.expectedString {
// 				t.Errorf("got %s, wanted %s", ans, tt.expectedString)
// 			}
// 		})
// 	}
// 	fmt.Println()
// }
//
