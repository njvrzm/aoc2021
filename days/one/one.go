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
	last := input[0]
	ups := 0
	for _, val := range input {
		if val > last {
			ups += 1
		}
		last = val
	}
	return ups
}

func PartTwo(input Input) int {
	ups := 0
	for i := range input {
		if i > 2 {
			if input[i] > input[i-3] {
				ups += 1
			}
		}
	}
	return ups
}
