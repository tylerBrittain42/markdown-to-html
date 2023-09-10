package main

import (
	"bufio"
	"fmt"
	"os"
)

func ioAndFmt() {
	fmt.Println("What is your name")
	var name string
	var name2 string
	_, _ = fmt.Scanln(&name, &name2)
	fmt.Println("Hello " + name2 + " " + name)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("done")

	file, e := os.Open("./read_input.md")
	check(e)

	// bytes := make([]byte, 100)
	// file.Read(bytes)
	// fmt.Println(bytes)
	// fmt.Println(string(bytes) + "end")

	scanner = bufio.NewScanner(file)
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
}
