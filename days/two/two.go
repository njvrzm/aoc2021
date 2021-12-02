package two

import (
	"aoc/help"
	"strconv"
	"strings"
)

type Input []string

func GetInput(path string) Input {
	return help.ReadInput(path)
}

func PartOne(input Input) int {
	depth, farth := 0, 0
	for _, line := range input {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		switch direction {
		case "forward":
			farth += distance
		case "down":
			depth += distance
		case "up":
			depth -= distance
		}
	}
	return depth * farth
}

func PartTwo(input Input) int {
	depth, farth, aim := 0, 0, 0
	for _, line := range input {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		switch direction {
		case "forward":
			farth += distance
			depth += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		}
	}
	return depth * farth
}
