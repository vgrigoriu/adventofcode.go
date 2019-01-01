package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/vgrigoriu/adventofcode.go/common"
)

func main() {
	lines := common.Input()
	sort.Strings(lines)
	shifts := shifts(lines)
	guardians := make(map[int]guardianHistory)
	for _, shift := range shifts {
		guard := guardians[shift.id]
		guardians[shift.id] = guard.mark(shift)
	}
	maxID, maxMinute := 0, 0
	var maxGuard guardianHistory
	for id, guardian := range guardians {
		min := guardian.minuteMostAsleep()
		times := guardian.minutes[min]
		if times > maxMinute {
			maxGuard = guardian
			maxMinute = times
			maxID = id
		}
	}
	fmt.Println(maxID * maxGuard.minuteMostAsleep())
}

func shifts(lines []string) []Shift {
	result := make([]Shift, 0)
	shiftLines := make([]string, 0)
	shiftLines = append(shiftLines, lines[0])
	for _, line := range lines[1:] {
		if strings.Contains(line, "Guard") {
			result = append(result, newShift(shiftLines))
			shiftLines = make([]string, 0)
		}
		shiftLines = append(shiftLines, line)
	}
	// append last shift
	result = append(result, newShift(shiftLines))
	return result
}

// Shift is a night shift.
type Shift struct {
	lines         []string
	id            int
	minutes       [60]bool
	minutesAsleep int
}

func newShift(lines []string) Shift {
	// get guard id
	var id int
	fmt.Sscanf(lines[0][19:], "Guard #%d begins shift", &id)
	// compute minutes asleep
	minutesAsleep := 0
	minutes := [60]bool{}
	for i := 1; i < len(lines); i += 2 {
		var start, end int
		fmt.Sscanf(lines[i][15:], "%v", &start)
		fmt.Sscanf(lines[i+1][15:], "%v", &end)
		for m := start; m < end; m++ {
			minutes[m] = true
		}
		minutesAsleep += (end - start)
	}
	return Shift{lines, id, minutes, minutesAsleep}
}

type guardianHistory struct {
	totalMinutes int
	minutes      [60]int
}

func (g guardianHistory) mark(shift Shift) guardianHistory {
	g.totalMinutes += shift.minutesAsleep
	for i := 0; i < 60; i++ {
		if shift.minutes[i] {
			g.minutes[i]++
		}
	}

	return g
}

func (g guardianHistory) minuteMostAsleep() int {
	maxMinute := 0
	for i := 0; i < 60; i++ {
		if g.minutes[i] >= g.minutes[maxMinute] {
			maxMinute = i
		}
	}
	return maxMinute
}
