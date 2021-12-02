package one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartOne(t *testing.T) {
	tt := []struct{
		Input Input
		Expected int
	}{
		{
			Input{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			7,
		},
		{
			GetInput("../../inputs/one"),
			1676,
		},
	}
	for _, testCase := range tt {
		assert.Equal(t, testCase.Expected, PartOne(testCase.Input))
	}
}


func TestPartTwo(t *testing.T) {
	tt := []struct{
		Input Input
		Expected int
	}{
		{
			Input{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			5,
		},
		{
			GetInput("../../inputs/one"),
			1706,
		},
	}
	for _, testCase := range tt {
		assert.Equal(t, testCase.Expected, PartTwo(testCase.Input))
	}
}

