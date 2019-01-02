// Package common implents stuff useful for more than one Advent of Code problem.
package common

import (
	"bufio"
	"flag"
	"log"
	"os"
)

// Input returns the content of the input files as a slice of strings.
func Input() []string {
	flag.Parse()

	args := flag.Args()
	if len(args) > 1 {
		log.Fatalf("don't know what to do with more than 1 argument: %v", args)
	}

	input := "input.txt"
	if len(args) == 1 {
		input = args[0]
	}
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// SingleLineInput returns the only line from the input file, or panics.
func SingleLineInput() string {
	input := Input()
	if len(input) != 1 {
		panic("expecting exactly one line of input")
	}
	return input[0]
}
