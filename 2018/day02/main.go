package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	by2, by3 := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if hasLetterTwice(line) {
			by2++
		}
		if hasLetterThrice(line) {
			by3++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(by2 * by3)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)
}

func hasLetterTwice(s string) bool {
	letters := make(map[rune]int)
	for _, r := range s {
		letters[r]++
	}

	for _, count := range letters {
		if count == 2 {
			return true
		}
	}

	return false
}

func hasLetterThrice(s string) bool {
	letters := make(map[rune]int)
	for _, r := range s {
		letters[r]++
	}

	for _, count := range letters {
		if count == 3 {
			return true
		}
	}

	return false
}

func differByOneLetter(a, b string) (bool, string) {
	diffs := 0
	common := []byte{}
	for i := 0; i < len(a); i++ {
		// ignore that runes can have more than one byte, assume that everything is ASCII
		// ignore that strings might have different lengths
		if a[i] != b[i] {
			diffs++
		} else {
			common = append(common, a[i])
		}
	}

	return diffs == 1, string(common)
}
