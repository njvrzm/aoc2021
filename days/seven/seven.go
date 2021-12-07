package seven

import (
	"aoc/help"
	"math"
	"strings"
)

func GetInput(s string) []int {
	out := []int{}
	for _, v := range strings.Split(s, ",") {
		out = append(out, help.Sinter(v))
	}
	return out
}

func Linear(i int) int {
	return i
}
func Triangular(i int) int {
	return (i*(i+1)) >> 1
}
// PartOne takes a list of horizontal positions and finds the position
// that all can be moved to with the least total displacement
func PartOne(crabs []int, costFunction func(int) int) (int, int) {
	low := help.Min(crabs)
	high := help.Max(crabs)
	cost := math.MaxInt
	cheapest := math.MaxInt

	for i := low; i < high; i++ {
		t := 0
		for _, v := range crabs {
			 t += costFunction(help.Abs(i - v))
		}
		if t < cost {
			cost = t
			cheapest = i
		}
	}
	return cost, cheapest
}