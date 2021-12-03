package three

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayThreePartOne(t *testing.T) {
	tt := []struct{
		Input []string
		Expected int
	}{
		{
			[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			198,
		},
		{
			GetInput("../../inputs/three"),
			3895776,
		},
	}
	for _, testCase := range tt {
		assert.Equal(t, testCase.Expected, PartOne(testCase.Input))
	}
}
func TestDayThreePartTwo(t *testing.T) {
	tt := []struct{
		Input []string
		Expected int
	}{
		{
			[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			230,
		},
		{
			GetInput("../../inputs/three"),
			7928162,
		},
	}
	for _, testCase := range tt {
		assert.Equal(t, testCase.Expected, PartTwo(testCase.Input))
	}
}
