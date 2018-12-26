package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
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
