package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBuildLine(t *testing.T) {
	tests := []struct {
		inputString   string
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
			// if ans.Equals(&tt.expectedStack){
			if ans != tt.expectedString {
				t.Errorf("got %s, wanted %s", ans, tt.expectedString)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	mockIO := []struct {
		testName string
		input string
		output string
	}{
		{"Single block", "# This is single line", "<h1>This is a single line</h1>"},
		{"Single block with bold and italic", "## **This** *is* single line", "<h2><strong>This</strong> <em>is</em> a single line</h2>"},
	}

	// var buffer bytes.Buffer
	// buffer.WriteString("This is a test\nThis is the second line of the test")
	// testScanner := bufio.NewScanner(&buffer)
	// testScanner.Scan()
	// fmt.Println(testScanner.Text())
	// fmt.Println("z")
	// testScanner.Scan()
	// fmt.Println(testScanner.Text())

	// TESTING STARTS HERE
	fmt.Println("Convert tests")
	
	// Note: This can technically be one loop, but I am focusing on readability

	// load buffer
	var readBuffer, writerBuffer bytes.Buffer
	for _, mock := range mockIO {
		readBuffer.WriteString(mock.input + "\n")	
	}

	//testing
	convert(&readBuffer, &writerBuffer)
	fmt.Println("here")
	fmt.Println(writerBuffer.String())
	fmt.Println("here")

	//perform check here

	// for _, mock := range mockIO {
	// 	testname := fmt.Sprintf("name: %s", mock.testName)
	// 	fmt.Println(testname)
	// 	t.Run(testname, func(t *testing.T) {
	// 		convert(tt.inputFile)
	//
	// 		if ans != tt.expectedStack {
	// 			t.Errorf("got %s, wanted %s", ans, tt.expectedStack)
	// 		}
	// 	})
	// }

	// fmt.Println("Convert tests")
	// for _, tt := range tests {
	// 	testname := fmt.Sprintf("name: %s", tt.testName)
	// 	fmt.Println(testname)
	// 	t.Run(testname, func(t *testing.T) {
	// 		convert(tt.inputFile)
	//
	// 		if ans != tt.expectedStack {
	// 			t.Errorf("got %s, wanted %s", ans, tt.expectedStack)
	// 		}
	// 	})
	// }
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
