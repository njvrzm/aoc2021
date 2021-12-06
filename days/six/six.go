package six

import (
	"aoc/help"
	"strings"
)



type School struct {
	FishCount map[int]int
}

func (e *School) FromString(s string) {
	e.FishCount = make(map[int]int)
	for _, v := range strings.Split(s, ",") {
		e.FishCount[help.Sinter(v)] += 1
	}
}

func (e *School) Tick() {
	newSix := e.FishCount[0]
	for i := 1; i <= 8; i++ {
		e.FishCount[i-1] = e.FishCount[i]
	}
	e.FishCount[6] += newSix
	e.FishCount[8] = newSix
}

func (e *School) Total() int {
	t := 0
	for _, v := range e.FishCount {
		t += v
	}
	return t
}