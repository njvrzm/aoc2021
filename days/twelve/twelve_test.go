package twelve

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayTwelvePartOne(t *testing.T) {
	testCases := []struct {
		path     string
		expected int
	}{
		{path: "testdata/example_one", expected: 10},
		{path: "testdata/example_two", expected: 19},
		{path: "testdata/example_three", expected: 226},
		{path: "../../inputs/twelve", expected: 4411},
	}
	for _, tt := range testCases {
		caves := GetInput(tt.path)
		count := caves.CountEgresses(caves.GetCave("start"), false, make(map[string]int))
		assert.Equal(t, tt.expected, count)
	}
}
func TestDayTwelvePartTwo(t *testing.T) {
	testCases := []struct {
		path     string
		expected int
	}{
		{path: "testdata/example_one", expected: 36},
		{path: "testdata/example_two", expected: 103},
		{path: "testdata/example_three", expected: 3509},
		{path: "../../inputs/twelve", expected: 136767},
	}
	for _, tt := range testCases {
		caves := GetInput(tt.path)
		count := caves.CountEgresses(caves.GetCave("start"), true, make(map[string]int))
		assert.Equal(t, tt.expected, count)
	}
}
