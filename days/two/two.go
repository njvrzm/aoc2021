package two

import (
	"aoc/help"
	"strings"
)

type Direction string

type Command struct {
	Direction string
	Distance int
}

type Input []Command

func GetInput(path string) []Command {
	lines := help.ReadInput(path)
	out := make([]Command, len(lines))
	for _, line := range lines {
		parts := strings.SplitN(line, " ", 2)
		out = append(out, Command{parts[0], help.Sinter(parts[1])})
	}
	return out
}

func PartOne(input Input) int {
	return solve(input, false)
}

func PartTwo(input Input) int {
	return solve(input, true)
}

func solve(input Input, partTwo bool) int {
	depth, farth, aim := 0, 0, 0
	for _, command := range input {
		fFactor, dFactor, aFactor := 0, 0, 0
		switch command.Direction {
		case "forward":
			fFactor = 1
			dFactor = aim
		case "down":
			if partTwo {
				aFactor = 1
			} else {
				dFactor = 1
			}
		case "up":
			if partTwo {
				aFactor = -1
			} else {
				dFactor = -1
			}
		}
		farth += fFactor * command.Distance
		depth += dFactor * command.Distance
		aim += aFactor * command.Distance
	}
	return depth * farth
}
