package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/tylerBrittain42/markdown-to-html/parser"
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

type firstPassTest struct {
	caseName   string
	inputFile  []string
	outputFile []string
}

func createFirstPassCases() []firstPassTest {
	caseName := []string{
		"1. Single block line",
		"2. Single paragraph line",
		"3. Single ordered list line",
		"4. Multiple block line",
		"5. Multiple block line with italics/bold",
		"6. Multiple block and paragraph line",
		"7. Multiple block and ordered list",
		"8. Multiple block, paragraph, and ordered list with italics and bold",
		"9. Final test case, full document",
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
			"1. This is the second element of a list",
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
			"- This is part of a list too",
		},
		// 8
		{
			"# This is the *first* line",
			"- This is part of a **list**",
			"This is part of a paragraph",
		},
		{
			"# Cookies or Ice Cream",
			"by *Tyler*",
			"## Introduction",
			"The topic of the  **best** dessert food is one that has never had a definitive answer. This is often because there are too many factors of taste that must be accounted for.",
			"",
			"These factors include the following:",
			"- taste",
			"- temperature",
			"- price",
			"",
			"Only by considering each of these factors can we truly determine the ultimate dessert.",
			"",
			"### Conclusion",
			"The rankings for best desserts are as follows:",
			"1. Cookies",
			"1. Ice Cream",
			"Thank you",
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
		}, {
			// 3
			"<html>",
			"<ol>",
			"\t<li>This is the first element of a list</li>",
			"\t<li>This is the second element of a list</li>",
			"</ol>",
			"</html>",
		}, {
			// 4
			"<html>",
			"<h1>This is the first line</h1>",
			"<h2>This is the second line</h2>",
			"</html>",
		}, {
			// 5
			"<html>",
			"<h1>This is the <em>first</em> line</h1>",
			"<h2>This is the <strong>second</strong> line</h2>",
			"</html>",
		}, {
			// 6
			"<html>",
			"<h1>This is the first line</h1>",
			"<p>This is part of a paragraph</p>",
			"</html>",
		}, {
			// 7
			"<html>",
			"<h1>This is the first line</h1>",
			"<ul>",
			"\t<li>This is part of a list</li>",
			"\t<li>This is part of a list too</li>",
			"</ul>",
			"</html>",
		}, {
			// 8
			"<html>",
			"<h1>This is the <em>first</em> line</h1>",
			"<ul>",
			"\t<li>This is part of a <strong>list</strong></li>",
			"</ul>",
			"<p>This is part of a paragraph</p>",
			"</html>",
		}, {
			"<html>",
			"<h1>Cookies or Ice Cream</h1>",
			"<p>by <em>Tyler</em></p>",
			"<h2>Introduction</h2>",
			"<p>The topic of the  <strong>best</strong> dessert food is one that has never had a definitive answer. This is often because there are too many factors of taste that must be accounted for.</p>",
			"<p></p>",
			"<p>These factors include the following:</p>",
			"<ul>",
			"\t<li>taste</li>",
			"\t<li>temperature</li>",
			"\t<li>price</li>",
			"</ul>",
			"<p></p>",
			"<p>Only by considering each of these factors can we truly determine the ultimate dessert.</p>",
			"<p></p>",
			"<h3>Conclusion</h3>",
			"<p>The rankings for best desserts are as follows:</p>",
			"<ol>",
			"\t<li>Cookies</li>",
			"\t<li>Ice Cream</li>",
			"</ol>",
			"<p>Thank you</p>",
			"</html>",
		},
	}
	tests := []firstPassTest{}
	for i := range caseName {
		tests = append(tests, firstPassTest{caseName: caseName[i], inputFile: inputCases[i], outputFile: expectedOutputs[i]})
	}
	return tests
}

func TestFirstPass(t *testing.T) {

	// Input and output are separate variables because the initial <html> tag will make the lines off

	tests := createFirstPassCases()

	fmt.Println("\nTest: TestFirstPass")
	for _, tt := range tests {
		isCorrect := true
		// load buffer
		var readBuffer, writerBuffer bytes.Buffer
		for _, line := range tt.inputFile {
			readBuffer.WriteString(string(line) + "\n")
		}

		testname := fmt.Sprintf("Case: %s", tt.caseName)
		// if tt.caseName == "9. Final test case, full document" {
		// 	continue
		// }
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			firstPass(&readBuffer, &writerBuffer)
			checkOutput := bufio.NewScanner(&writerBuffer)
			for j, mock := range tt.outputFile {
				checkOutput.Scan()
				if string(mock) != checkOutput.Text() {
					t.Errorf("%v) got .%s., wanted .%s.", j, checkOutput.Text(), string(mock))
					isCorrect = false
					break
				}
			}
		})
		// TODO: FIGURE OUTo WHY THE BUFFER NEEDS TO BE READ AGAIN
		// Look into Tee reader
		if !isCorrect {
			for _, line := range tt.inputFile {
				readBuffer.WriteString(string(line) + "\n")
			}

			firstPass(&readBuffer, &writerBuffer)
			firstPassOutput := bufio.NewScanner(&writerBuffer)
			fmt.Println("BEGIN FAILED OUTPUT")
			for firstPassOutput.Scan() {
				fmt.Println(firstPassOutput.Text())
			}
			isCorrect = true
			fmt.Println("END FAILED OUTPUT")
		}
	}

}

func createSecondPassCases() []firstPassTest {
	caseName := []string{
		"1. No paragraphs at all",
		"2. No empty paragraphs(nonempty paragraps exist)",
		"3. Only empty paragraph",
		"4. Final test case, full document",
	}
	inputCases := [][]string{
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
		}, {
			// 3
			"<html>",
			"<p></p>",
			"</html>",
		}, {
			// 4
			"<html>",
			"<h1>Cookies or Ice Cream</h1>",
			"<p>by <em>Tyler</em></p>",
			"<h2>Introduction</h2>",
			"<p>The topic of the  <strong>best</strong> dessert food is one that has never had a definitive answer. This is often because there are too many factors of taste that must be accounted for.</p>",
			"<p></p>",
			"<p>These factors include the following:</p>",
			"<ul>",
			"\t<li>taste</li>",
			"\t<li>temperature</li>",
			"\t<li>price</li>",
			"</ul>",
			"<p></p>",
			"<p>Only by considering each of these factors can we truly determine the ultimate dessert.</p>",
			"<p></p>",
			"<h3>Conclusion</h3>",
			"<p>The rankings for best desserts are as follows:</p>",
			"<ol>",
			"\t<li>Cookies</li>",
			"\t<li>Ice Cream</li>",
			"</ol>",
			"<p>Thank you</p>",
			"</html>",
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
		}, {
			// 3
			"<html>",
			"</html>",
		}, {
			// 4
			"<html>",
			"<h1>Cookies or Ice Cream</h1>",
			"<p>by <em>Tyler</em></p>",
			"<h2>Introduction</h2>",
			"<p>The topic of the  <strong>best</strong> dessert food is one that has never had a definitive answer. This is often because there are too many factors of taste that must be accounted for.</p>",
			"<p>These factors include the following:</p>",
			"<ul>",
			"\t<li>taste</li>",
			"\t<li>temperature</li>",
			"\t<li>price</li>",
			"</ul>",
			"<p>Only by considering each of these factors can we truly determine the ultimate dessert.</p>",
			"<h3>Conclusion</h3>",
			"<p>The rankings for best desserts are as follows:</p>",
			"<ol>",
			"\t<li>Cookies</li>",
			"\t<li>Ice Cream</li>",
			"</ol>",
			"<p>Thank you</p>",
			"</html>",
		},
	}
	tests := []firstPassTest{}
	for i := range caseName {
		tests = append(tests, firstPassTest{caseName: caseName[i], inputFile: inputCases[i], outputFile: expectedOutputs[i]})
	}
	return tests
}

func TestSecondPass(t *testing.T) {

	// Input and output are separate variables because the initial <html> tag will make the lines off

	tests := createSecondPassCases()

	fmt.Println("\nTest: TestSecondPass")
	for _, tt := range tests {
		isCorrect := true
		// load buffer
		var readBuffer, writerBuffer bytes.Buffer
		for _, line := range tt.inputFile {
			readBuffer.WriteString(string(line) + "\n")
		}

		testname := fmt.Sprintf("Case: %s", tt.caseName)
		// if tt.caseName == "9. Final test case, full document" {
		// 	continue
		// }
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			secondPass(&readBuffer, &writerBuffer)
			checkOutput := bufio.NewScanner(&writerBuffer)
			for j, mock := range tt.outputFile {
				checkOutput.Scan()
				if string(mock) != checkOutput.Text() {
					t.Errorf("%v) got .%s., wanted .%s.", j, checkOutput.Text(), string(mock))
					isCorrect = false
					break
				}
			}
		})
		// TODO: FIGURE OUTo WHY THE BUFFER NEEDS TO BE READ AGAIN
		// Look into Tee reader
		if !isCorrect {
			for _, line := range tt.inputFile {
				readBuffer.WriteString(string(line) + "\n")
			}

			firstPass(&readBuffer, &writerBuffer)
			firstPassOutput := bufio.NewScanner(&writerBuffer)
			fmt.Println("BEGIN FAILED OUTPUT")
			for firstPassOutput.Scan() {
				fmt.Println(firstPassOutput.Text())
			}
			isCorrect = true
			fmt.Println("END FAILED OUTPUT")
		}
	}

}



// func TestConvertDocument(t *testing.T) {
//
// 	inputFile := "/test/final_input.md"
// 	correctOutputFile := "/test/final_output.html"
// 	convertOutputFile := "/test/test_output.html"
//
// 	fmt.Println("Test: TestConvert")
// 	for _, tt := range tests {
// 		isCorrect := true
// 		// load buffer
// 		var readBuffer, writerBuffer bytes.Buffer
// 		for _, line := range tt.inputFile {
// 			readBuffer.WriteString(string(line) + "\n")
// 		}
//
// 		testname := fmt.Sprintf("Case: %s", tt.caseName)
// 		fmt.Println(testname)
// 		t.Run(testname, func(t *testing.T) {
// 			convert(&readBuffer, &writerBuffer)
// 			checkOutput := bufio.NewScanner(&writerBuffer)
// 			for j, mock := range tt.outputFile {
// 				checkOutput.Scan()
// 				if string(mock) != checkOutput.Text() {
// 					t.Errorf("%v) got .%s., wanted .%s.", j, checkOutput.Text(), string(mock))
// 					isCorrect = false
// 					break
// 				}
// 			}
// 		}
//
// 	}
// }
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
