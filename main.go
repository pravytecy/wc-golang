/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type flagOptions struct {
	lines bool
	words bool
	chars bool
}

var (
	flagSet    flagOptions
	linesCount int
	wordCount  int
	charCount  int
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		printToStderr(err)
	}

}

func init() {
	rootCmd.Flags().BoolVarP(&flagSet.lines, "lines", "l", false, "Count number of lines")
	rootCmd.Flags().BoolVarP(&flagSet.words, "words", "w", false, "Count number of words")
	rootCmd.Flags().BoolVarP(&flagSet.chars, "chars", "c", false, "Count number of characters")

}

var rootCmd = &cobra.Command{
	Use:   "wc",
	Short: "wc is a word,line and length count tool similiar to unix wc command ",
	Long:  `wc is a word, line, and character count tool that reads from the standard input or from a file and outputs the number of lines, words, and characters`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			args = []string{"-"}
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if scanner.Err() != nil {
			// Handle error.
		}
	},
}

func printToStderr(err error) {
	fmt.Fprint(os.Stderr, err.Error())
}
