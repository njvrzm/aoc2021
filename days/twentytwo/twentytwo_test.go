package twentytwo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayTwentyTwo_Partial(t *testing.T) {
	testCases := []struct {
		cubes       Cuboids
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
		assert.Equal(t, tt.expected_on, tt.cubes.Union(true))
	}

}
func TestDayTwentyTwo_Full(t *testing.T) {
	testCases := []struct {
		cubes       Cuboids
		expected_on int
	}{
		{
			GetInput("testdata/example_three"),
			2758514936282235,
		},
		{
			GetInput("../../inputs/twentytwo"),
			1294137045134837,
		},
	}
	for _, tt := range testCases {
		assert.Equal(t, tt.expected_on, tt.cubes.Union(false))
	}

}
