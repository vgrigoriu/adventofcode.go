package main

import (
	"fmt"
	"github.com/vgrigoriu/adventofcode.go/common"
	"sort"
	"strings"
)

func main() {
	lines := common.Input()
	sort.Strings(lines)
	shifts := shifts(lines)
	fmt.Println(shifts)
}

func shifts(lines []string) []Shift {
	result := make([]Shift, 0)
	shift := make(Shift, 0)
	shift = append(shift, lines[0])
	for _, line := range lines[1:] {
		if strings.Contains(line, "Guard") {
			result = append(result, shift)
			shift = make(Shift, 0)
		}
		shift = append(shift, line)
	}
	// append last shift
	result = append(result, shift)
	return result
}

// Shift is a night shift.
type Shift []string
