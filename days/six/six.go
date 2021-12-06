package six

import (
	"aoc/help"
	"strings"
)



type School struct {
	FishCount []int
}

func (e *School) FromString(s string) {
	e.FishCount = make([]int, 9, 9)
	for _, v := range strings.Split(s, ",") {
		e.FishCount[help.Sinter(v)] += 1
	}
}

func (e *School) Tick() {
	e.FishCount = append(e.FishCount[1:], e.FishCount[0])
	e.FishCount[6] += e.FishCount[8]
}

func (e *School) Total() int {
	t := 0
	for _, v := range e.FishCount {
		t += v
	}
	return t
}