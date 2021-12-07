package seven

import (
	"aoc/help"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDaySevenPartOne_Example(t *testing.T) {
	input := GetInput("16,1,2,0,4,2,7,1,2,14")

	cost, cheapest := PartOne(input, Linear)
	assert.Equal(t, 37, cost)
	assert.Equal(t, 2, cheapest)
}
func TestDaySevenPartOne_Input(t *testing.T) {
	input := GetInput(help.ReadInput("../../inputs/seven")[0])

	cost, cheapest := PartOne(input, Linear)
	assert.Equal(t, 364898, cost)
	assert.Equal(t, 361, cheapest)
}
func TestDaySevenPartTwo_Example(t *testing.T) {
	input := GetInput("16,1,2,0,4,2,7,1,2,14")

	cost, cheapest := PartOne(input, Triangular)
	assert.Equal(t, 168, cost)
	assert.Equal(t, 5, cheapest)
}
func TestDaySevenPartTwo_Input(t *testing.T) {
	input := GetInput(help.ReadInput("../../inputs/seven")[0])

	cost, cheapest := PartOne(input, Triangular)
	assert.Equal(t, 104149091, cost)
	assert.Equal(t, 500, cheapest)
}
