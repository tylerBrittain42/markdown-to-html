package main

import (
	"fmt"
	"testing"
)

func TestHtmlBuilder(t *testing.T) {
	tests := []struct {
		inputBlock        string
		inputContent string
		expectedHtml string
	}{
		{"h1","Title", "<h1>Title</h1>"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s, %s", tt.inputBlock, tt.inputContent)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := htmlBuilder(tt.inputBlock,tt.inputContent)
			if ans != tt.expectedHtml {
				t.Errorf("got %s, wanted %s", ans, tt.expectedHtml)
			}
		})
	}
}
func TestGetContent(t *testing.T) {
	tests := []struct {
		input        string
		expectedContent string
	}{
		{"# singleSpace", "singleSpace"},
		{"# this has multiple spaces", "this has multiple spaces"},
		// {"#noSpaces", "<h3>"}, add for error handling
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.input)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := getContents(tt.input)
			if ans != tt.expectedContent {
				t.Errorf("got %s, wanted %s", ans, tt.expectedContent)
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
		{"--- this is a line break", "/br"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.input)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := getBlockType(tt.input)
			if ans != tt.expectedOpen {
				t.Errorf("got %s, wanted %s", ans, tt.expectedOpen)
			}
		})
	}

}
