package eight

import (
	"aoc/help"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDayEightPartOne_Example(t *testing.T) {
	entries := GetInput("testdata/example")
	assert.Equal(t, 26, PartOne(entries))
}
func TestDayEightPartOne_Input(t *testing.T) {
	entries := GetInput("../../inputs/eight")
	assert.Equal(t, 397, PartOne(entries))
}
func TestDayEightPartTwo_Example(t *testing.T) {
	entries := GetInput("testdata/example")
	total := 0
	for _, entry := range entries {
		eton := Identify(entry.Signals)
		out := []string{}
		for _, o := range entry.Outputs {
			ols := NewLightSet(o)
			for s, i := range eton {
				if ols.Equals(NewLightSet(s)) {
					out = append(out, fmt.Sprint(i))
				}
			}
		}
		total += help.Sinter(strings.Join(out, ""))
	}
	assert.Equal(t, 61229, total)
}
func TestDayEightPartTwo_Input(t *testing.T) {
	entries := GetInput("../../inputs/eight")
	total := 0
	for _, entry := range entries {
		eton := Identify(entry.Signals)
		out := []string{}
		for _, o := range entry.Outputs {
			ols := NewLightSet(o)
			for s, i := range eton {
				if ols.Equals(NewLightSet(s)) {
					out = append(out, fmt.Sprint(i))
				}
			}
		}
		total += help.Sinter(strings.Join(out, ""))
	}
	assert.Equal(t, 1027422, total)
}
