package ten

import (
	"aoc/help"
	"sort"
	"strings"
)

var IncorrectScores = map[string]int {
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,

}
var IncompleteScores = map[string]int {
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}
const Open = "([{<"
const Close = ")]}>"

func GetInput(path string) []string {
	return help.ReadInput(path)
}

func Parse(line string) ([]string, string) {
	stack := []string{}
	for _, r := range line {
		rs := string(r)
		if place := strings.Index(Open, rs); place >= 0 {
			stack = append(stack, rs)
		} else if place := strings.Index(Close, rs); place >= 0{
			opener := string(Open[place])
			if stack[len(stack)-1] == opener {
				stack = stack[:len(stack)-1]
			} else {
				return stack, rs
			}
		}
	}
	return stack, ""
}

func IncorrectScore(lines []string) int {
	score := 0
	for _, line := range lines {
		_, incorrect := Parse(line)
		score += IncorrectScores[incorrect]
	}
	return score
}

func Complete(stack []string) string {
	out := []string{}
	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out = append(out, string(Close[strings.Index(Open, last)]))
	}
	return strings.Join(out, "")
}

func IncompleteScore(lines []string) int {
	scores := []int{}
	for _, line := range lines {
		stack, incorrect := Parse(line)
		if incorrect == "" {
			score := 0
			complete := Complete(stack)
			for _, r := range complete {
				score *= 5
				score += IncompleteScores[string(r)]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}