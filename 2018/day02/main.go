package main

import "fmt"

func main() {
	fmt.Println("Ana are mere")
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
