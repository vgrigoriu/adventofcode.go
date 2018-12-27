package main

import (
	"fmt"

	"github.com/vgrigoriu/adventofcode.go/common"
)

func main() {
	part1()
	part2()
}

func part1() {
	input := common.Input()

	by2, by3 := 0, 0
	for _, line := range input {
		if hasLetterTwice(line) {
			by2++
		}
		if hasLetterThrice(line) {
			by3++
		}
	}

	fmt.Println(by2 * by3)
}

func part2() {
	lines := common.Input()

	for i, id1 := range lines {
		for _, id2 := range lines[i+1:] {
			if ok, common := differByOneLetter(id1, id2); ok {
				fmt.Println(common)
			}
		}
	}
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
