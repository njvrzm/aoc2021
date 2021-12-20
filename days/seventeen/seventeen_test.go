package seventeen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDaySeventeenPartOne(t *testing.T) {
	testCases := []struct {
		box    Box
		height int
	}{
		{
			box:    GetInput("testdata/example"),
			height: 45,
		},
		{
			box:    GetInput("../../inputs/seventeen"),
			height: 19503,
		},
	}
	for _, tt := range testCases {
		height := PartOne(tt.box)
		assert.Equal(t, tt.height, height)
	}
}
func TestDaySeventeenPartTwo(t *testing.T) {
	testCases := []struct {
		box     Box
		vectors int
	}{
		{
			box:     GetInput("testdata/example"),
			vectors: 112,
		},
		{
			box:     GetInput("../../inputs/seventeen"),
			vectors: 5200,
		},
	}
	for _, tt := range testCases {
		height := PartTwo(tt.box)
		assert.Equal(t, tt.vectors, height)
	}
}
