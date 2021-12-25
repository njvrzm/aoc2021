package twentyfive

import (
	"aoc/help"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayTwentyFive(t *testing.T) {
	testCases := []struct {
		path          string
		expectedSteps int
	}{
		{
			"testdata/example",
			58,
		},
		{
			"../../inputs/twentyfive",
			571,
		},
	}
	for _, tt := range testCases {
		w := World{}
		w.FromLines(help.ReadInput(tt.path))
		for w.Step() {
		}
		assert.Equal(t, tt.expectedSteps, w.Steps)
	}
}
