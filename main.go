/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes   int
	printUsage bool
	name       string
}

var usageString = fmt.Sprintf(`Usage: %s  [-h|--help] A greeter application which prints the name you entered  number of times.`, os.Args[0])
var errPosArgSpecified = errors.New("positional args specifed")

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "%s", usageString)
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		if errors.Is(err, errPosArgSpecified) {
			fmt.Fprintln(os.Stdout, err)
		}
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}
	return nil
}

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("Jerry printer", flag.ContinueOnError)
	fs.SetOutput(w)

	fs.Usage = func() {
		var usageString = `A greeter application which prints the name you entered a specified number of times. Usage of %s:  [name]`
		fmt.Fprintf(w, usageString, fs.Name())
		fmt.Fprintln(w)

		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()

	}
	fs.IntVar(&c.numTimes, "n", 1, "Number of times to greet")
	fs.StringVar(&c.name, "a", "", "Welcome")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() > 1 {
		return c, errPosArgSpecified
	}
	if fs.NArg() == 1 {
		c.name = fs.Arg(0)
	}
	return c, nil
}

// func parseArgs(args []string) (config, error) {
// 	var numTimes int
// 	var err error

// 	c := config{}
// 	if len(args) != 1 {
// 		return c, errors.New("invalid number of arguments")
// 	}
// 	if args[0] == "-h" || args[0] == "--help" {
// 		c.printUsage = true
// 		return c, nil
// 	}
// 	if args[0] == "-l" {
// 		c.printUsage = true
// 		return c, nil
// 	}
// 	numTimes, err = strconv.Atoi(args[0])
// 	if err != nil {
// 		return c, err
// 	}
// 	c.numTimes = numTimes
// 	return c, nil
// }

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
	fmt.Fprintf(w, "%s", msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("you didn't enter your name")
	}
	return name, nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	var err error
	if len(c.name) == 0 {
		c.name, err = getName(r, w)
		if err != nil {
			return err
		}
	}
	greetUser(c, w)
	return nil
}

func greetUser(c config, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", c.name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}
