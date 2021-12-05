package five

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayFivePartOne_Example(t *testing.T) {
	segs := GetInput("testdata/five_example.txt")
	assert.Equal(t, 10, len(segs))
	assert.Equal(t, 5, Overlaps(segs, false))
}
func TestDayFivePartOne_Input(t *testing.T) {
	segs := GetInput("../../inputs/five")
	assert.Equal(t, 5145, Overlaps(segs, false))
}
func TestDayFivePartTwo_Example(t *testing.T) {
	segs := GetInput("testdata/five_example.txt")
	assert.Equal(t, 12, Overlaps(segs, true))
}
func TestDayFivePartTwo_Input(t *testing.T) {
	segs := GetInput("../../inputs/five")
	assert.Equal(t, 16518, Overlaps(segs, true))
}
