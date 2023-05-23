/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type flagOptions struct {
	lineFlag bool
	wordFlag bool
	charFlag bool
}

var (
	flagSet                                        flagOptions
	totalLineCount, totalWordCount, totalCharCount int
)

// rootCmd represents the base command when called without any subcommands
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wc-golang.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&flagSet.lineFlag, "lines", "l", false, "Count number of lines")
}
