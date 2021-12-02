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
			Input{
				Command{"forward", 5},
				Command{"down",  5},
				Command{"forward",  8},
				Command{"up",  3},
				Command{"down",  8},
				Command{"forward",  2},
			},
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
			Input{
				Command{"forward", 5},
				Command{"down",  5},
				Command{"forward",  8},
				Command{"up",  3},
				Command{"down",  8},
				Command{"forward",  2},
			},
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

