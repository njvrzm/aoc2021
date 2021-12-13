package eleven

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayElevenPartOne_Example(t *testing.T) {
	field := GetInput("testdata/example")
	assert.Equal(t, 1656, field.ManyStep(100))
}
func TestDayElevenPartOne_Input(t *testing.T) {
	field := GetInput("../../inputs/eleven")
	assert.Equal(t, 1785, field.ManyStep(100))
}
func TestDayElevenPartTwo_Example(t *testing.T) {
	field := GetInput("testdata/example")
	assert.Equal(t, 195, field.Seek())
}
func TestDayElevenPartTwo_Input(t *testing.T) {
	field := GetInput("../../inputs/eleven")
	assert.Equal(t, 354, field.Seek())
}
