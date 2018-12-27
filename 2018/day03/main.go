package main

import (
	"fmt"
	"log"

	"github.com/vgrigoriu/adventofcode.go/common"
)

var fabric [1000][1000]int

func main() {
	lines := common.Input()

	claims := make([]claim, len(lines))
	for i, line := range lines {
		claims[i] = new(line)
	}

	for _, claim := range claims {
		claim.stake()
	}

	conflicts := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] > 1 {
				conflicts++
			}
		}
	}

	fmt.Println(conflicts)
}

type claim struct {
	id, left, top, w, h int
}

func new(s string) claim {
	// #123 @ 3,2: 5x4
	var id, left, top, w, h int
	_, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &left, &top, &w, &h)
	if err != nil {
		log.Fatalf("cannot parse a claim out of «%s»", s)
	}
	return claim{id, left, top, w, h}
}

func (c claim) stake() {
	for i := c.left; i < c.left+c.w; i++ {
		for j := c.top; j < c.top+c.h; j++ {
			fabric[i][j]++
		}
	}
}

func (c claim) String() string {
	return fmt.Sprintf("id: %d; left: %d; top: %d; w: %d; w: %d", c.id, c.left, c.top, c.w, c.h)
}
