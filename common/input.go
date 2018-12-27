// Package common implents stuff useful for more than one Advent of Code problem.
package common

import (
	"bufio"
	"log"
	"os"
)

// Input returns the content of the input files as a slice of strings.
func Input() []string {
	file, err := os.Open("input.txt")
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
