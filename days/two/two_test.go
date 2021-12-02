package two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayTwoPartOne(t *testing.T) {
	tt := []struct{
		Input Input
		Expected int
	}{
		{
			Input{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			150,
		},
		{
			GetInput("../../inputs/two"),
			1938402,
		},
	}
	for _, testCase := range tt {
		assert.Equal(t, testCase.Expected, PartOne(testCase.Input))
	}
}


func TestDayTwoPartTwo(t *testing.T) {
	tt := []struct{
		Input Input
		Expected int
	}{
		{
			Input{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			900,
		},
		{
			GetInput("../../inputs/two"),
			1947878632,
		},
	}
	for _, testCase := range tt {
		assert.Equal(t, testCase.Expected, PartTwo(testCase.Input))
	}
}

