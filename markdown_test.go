package main

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

func TestBuildLine(t *testing.T) {
	tests := []struct {
		inputString    string
		expectedString string
	}{
		{"# This is a single line", "<h1>This is a single line</h1>"},
		{"## This is a single line", "<h2>This is a single line</h2>"},
		{"### This has **underlined** characters", "<h3>This has <strong>underlined</strong> characters</h3>"},
		{"## This has **underlined** characters and *italic ones*", "<h2>This has <strong>underlined</strong> characters and <em>italic ones</em></h2>"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := buildLine(tt.inputString)
			if ans != tt.expectedString {
				t.Errorf("got %s, wanted %s", ans, tt.expectedString)
			}
		})
	}
	fmt.Println()
}

type convertCase struct {
	testName string
	input    string
	output   string
}

func TestConvert(t *testing.T) {

	tests := [][]convertCase{
		{
			{"Single block", "# This is a single line", "<h1>This is a single line</h1>"},
			{"Single block with bold and italic", "## **This** *is* a single line", "<h2><strong>This</strong> <em>is</em> a single line</h2>"},
		},
		{
			{"Single block2", "### This is a single line from the second case", "<h3>This is a single line from the second case</h3>"},
			{"Single block with bold and italic", "## **This** *is* a single line", "<h2><strong>This</strong> <em>is</em> a single line</h2>"},
		},
	}

	// TESTING STARTS HERE
	// Note: This can technically be one loop, but I am focusing on readability

	for _, tt := range tests {

		// load buffer
		var readBuffer, writerBuffer bytes.Buffer
		for _, mock := range tt {
			readBuffer.WriteString(mock.input + "\n")
		}

		//perform check here
		testname := fmt.Sprintf("name: %s", "Convert() test")
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			checkOutput := bufio.NewScanner(&writerBuffer)
			convert(&readBuffer, &writerBuffer)
			checkOutput.Scan()
			for _, mock := range tt {
				if mock.output != checkOutput.Text() {
					t.Errorf("got .%s., wanted .%s.", checkOutput.Text(), mock.output)
					break
				}
				checkOutput.Scan()
			}
		})
	}

}

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
