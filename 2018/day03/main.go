package main

import (
	"fmt"

	"github.com/vgrigoriu/adventofcode.go/common"
)

func main() {
	lines := common.Input()

	for _, line := range lines {
		fmt.Println(line)
	}
}

type claim struct {
	left, top, w, h int
}

func new(s string) claim {
	// #123 @ 3,2: 5x4
	var id, left, top, w, h int
	fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &left, &top, &w, &h)
	return claim{left, top, w, h}
}
