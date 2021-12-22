package twentyone

import (
	"aoc/help"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayTwentyOne(t *testing.T) {
	testCases := []struct {
		playerOne int
		playerTwo int
		expected  int
	}{
		{4, 8, 739785},
		{5, 9, 989352},
	}
	for _, tt := range testCases {
		assert.Equal(t, tt.expected, PartOne(tt.playerOne, tt.playerTwo))
	}
}
func TestDayTwentyOne_PartTwo(t *testing.T) {
	testCases := []struct {
		playerOne int
		playerTwo int
		expected  int
	}{
		{4, 8, 444356092776315},
		{5, 9, 430229563871565},
	}
	for _, tt := range testCases {
		r := Play(tt.playerOne, tt.playerTwo)
		assert.Equal(t, tt.expected, help.Max(r.One, r.Two))
		fmt.Println(r.One + r.Two)
	}
}
