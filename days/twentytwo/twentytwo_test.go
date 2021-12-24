package twentytwo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayTwentyTwo(t *testing.T) {
	testCases := []struct {
		cubes       Cubes
		expected_on int
	}{
		{
			GetInput("testdata/example_one"),
			39,
		},
		{
			GetInput("testdata/example_two"),
			590784,
		},
		{
			GetInput("../../inputs/twentytwo"),
			583636,
		},
	}
	for _, tt := range testCases {
		assert.Equal(t, tt.expected_on, tt.cubes.EvaluateSlow())
	}

}
