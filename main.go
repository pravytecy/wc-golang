/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type flagOptins struct {
	lines bool
	words bool
	chars bool
}

var (
	flagSet    flagOptins
	linesCount int
	wordCount  int
)

func main() {

	lines := flag.Bool("l", false, "Number of lines")
	word := flag.Bool("w", false, "Number of words")
	flag.Parse()
	callOptions()
	otp := printValues(lines, word)

	printToStdout(otp)
}

// func parseArgs(w io.Writer, args []string) {
// 	//var c = &flagOptins{false, false, false}

// 	flagSet.lines = flag.Bool("l", false, "Number of lines")
// 	flag.Bool("w", false, "Number of words")
// 	//flag.BoolVar(&c.chars, "c", false, "Number of chars")
// 	flag.Parse()

// }

func callOptions() {
	countLines("sample.txt")
	countWords("sample.txt")
}

func countLines(file string) {
	filePath, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(filePath)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	filePath.Close()
	for _, line := range fileLines {
		fmt.Println(line)
		linesCount++
	}
	fmt.Printf("Number of lines : %d", linesCount)
}

func countWords(file string) {
	filePath, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(filePath)
	fileScanner.Split(bufio.ScanWords)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	filePath.Close()
	for _, line := range fileLines {
		fmt.Println(line)
		wordCount++
	}
}
func printValues(lines, words *bool) string {
	var output string
	// append only if lineFlag is set
	if *lines {
		output += fmt.Sprintf("%8d", linesCount)
	}

	// append only if wordFlag is set
	if *words {
		output += fmt.Sprintf("%8d", wordCount)
	}

	if !*lines && !*words {
		output += fmt.Sprintf("%8d", linesCount)
		output += fmt.Sprintf("%8d", wordCount)
	}
	return output
}

func printToStdout(s string) {
	fmt.Fprint(os.Stdout, s)
}
