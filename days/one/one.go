package one

import (
	"aoc/help"
	"strconv"
)

type Input []int

func GetInput(path string) Input {
	out := Input{}
	for _, line := range help.ReadInput(path) {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return out
}

func PartOne(input Input) int {
	return CountUpwardSpans(input, 1)
}

func PartTwo(input Input) int {
	return CountUpwardSpans(input, 3)
}

func CountUpwardSpans(input Input, span int) int {
	ups := 0
	for i := span; i < len(input); i++ {
		if input[i] > input[i-span] {
			 ups += 1
		}
	}
	return ups
}