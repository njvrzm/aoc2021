package nine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayNinePartOne_Example(t *testing.T) {
	f := GetInput("testdata/example")
	assert.Equal(t, 15, f.DangerTotal())
}
func TestDayNinePartOne_Input(t *testing.T) {
	f := GetInput("../../inputs/nine")
	assert.Equal(t, 489, f.DangerTotal())
}
func TestDayNinePartTwo_Example(t *testing.T) {
	f := GetInput("testdata/example")
	assert.Equal(t, 1134, f.BasinScore())
}
func TestDayNinePartTwo_Input(t *testing.T) {
	f := GetInput("../../inputs/nine")
	assert.Equal(t, 1056330, f.BasinScore())
}
