package six

import (
	"aoc/help"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDaySixPartOne_Example(t *testing.T) {
	input := "3,4,3,1,2"
	school := School{}
	school.FromString(input)
	for i := 0; i < 80; i ++ {
		school.Tick()
	}
	assert.Equal(t, 5934, school.Total())
}

func TestDaySixPartOne_Input(t *testing.T) {
	input := help.ReadInput("../../inputs/six")[0]
	school := School{}
	school.FromString(input)
	for i := 0; i < 80; i ++ {
		school.Tick()
	}
	assert.Equal(t, 386640, school.Total())
}

func TestDaySixPartTwo_Example(t *testing.T) {
	input := "3,4,3,1,2"
	school := School{}
	school.FromString(input)
	fmt.Println(school)
	for i := 0; i < 256; i ++ {
		school.Tick()
	}
	assert.Equal(t, 26984457539, school.Total())
}
func TestDaySixPartTwo_Input(t *testing.T) {
	input := help.ReadInput("../../inputs/six")[0]
	school := School{}
	school.FromString(input)
	fmt.Println(school)
	for i := 0; i < 256; i ++ {
		school.Tick()
	}
	assert.Equal(t, 1733403626279, school.Total())
}
