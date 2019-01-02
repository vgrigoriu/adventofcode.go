package main

import (
	"fmt"
	"github.com/vgrigoriu/adventofcode.go/common"
)

func main() {
	input := common.SingleLineInput()
	result, more := react(input)
	for more {
		result, more = react(result)
	}

	fmt.Println(len(result))
}

func react(input string) (result string, more bool) {
	for i := range input[:len(input)-1] {
		r1, r2 := rune(input[i]), rune(input[i+1])
		if r1-r2 == 32 || r2-r1 == 32 {
			return input[:i] + input[i+2:], true
		}
	}
	return input, false
}
