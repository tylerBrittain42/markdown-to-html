package main

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
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
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			actualBlock, actualBody := splitLine(tt.inputString)
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
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := getInnerText(tt.inputString)
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
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := openTag(tt.inputString)
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
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := closeTag(tt.inputString)
			if ans != tt.expectedString {
				t.Errorf("got .%s., wanted .%s.", ans, tt.expectedString)
			}
		})
	}
}

type convertTest struct {
	inputFile [][]string
	outputFile [][]string
}


func TestConvert(t *testing.T) {

	// Input and output are separate variables because the initial <html> tag will make the lines off

	inputCases := [][]string{
		{	
			"# This is a single line", 
			// "## **This** *is* a single line",
		},	
	}
	expectedOutputs := [][]string{
			{
			"<html>",
			"<h1>This is a single line</h1>",
			// "<h2><strong>This</strong> <em>is</em> a single line</h2>",
			"</html>",
			},
		}
	tests := []convertTest{{inputFile:inputCases, outputFile:expectedOutputs}}
		
				for i, tt := range tests {

		// load buffer
		var readBuffer, writerBuffer bytes.Buffer
		fmt.Println("Adding to readBuffer")
		for _, line := range tt.inputFile[i] {
			fmt.Println(line)
			readBuffer.WriteString(line + "\n")
		}

		//perform check here
		testname := fmt.Sprintf("name: %s", "Convert() test")
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			convert(&readBuffer, &writerBuffer)
			checkOutput := bufio.NewScanner(&writerBuffer)
			for _, mock := range tt.outputFile[i] {
				checkOutput.Scan()
				if mock != checkOutput.Text() {
					t.Errorf("got .%s., wanted .%s.", checkOutput.Text(), mock)
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
