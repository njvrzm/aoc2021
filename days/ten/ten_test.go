package ten

import (
	"github.com/stretchr/testify/assert"
	"testing"

	//"testing"
)

func TestDayTenPartOne_Example(t *testing.T) {
	lines := GetInput("testdata/example")
	assert.Equal(t, 26397, IncorrectScore(lines))
}
func TestDayTenPartOne_Input(t *testing.T) {
	lines := GetInput("../../inputs/ten")
	assert.Equal(t, 323613, IncorrectScore(lines))
}
func TestDayTenPartTwo_Example(t *testing.T) {
	lines := GetInput("testdata/example")
	assert.Equal(t, 288957, IncompleteScore(lines))
}
func TestDayTenPartTwo_Input(t *testing.T) {
	lines := GetInput("../../inputs/ten")
	assert.Equal(t, 3103006161, IncompleteScore(lines))
}
