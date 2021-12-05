package four

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayFourPartOne_Example(t *testing.T) {
	bingo := GetInput("testdata/four_example.txt")
	assert.Equal(t, 3, len(bingo.Boards))
	for _, board := range bingo.Boards {
		assert.False(t, board.Winning())
	}
	score := 0
	for {
		score = bingo.Call()
		if score > 0 {
			break
		}
	}
	assert.Equal(t, 4512, score)
}

func TestDayFourPartOne_Input(t *testing.T) {
	bingo := GetInput("../../inputs/four")
	score := 0
	for {
		score = bingo.Call()
		if score > 0 {
			break
		}
	}
	assert.Equal(t, 51034, score)
}
func TestDayFourPartTwo_Input(t *testing.T) {
	bingo := GetInput("../../inputs/four")
	score := 0
	for {
		score = bingo.CallMisere()
		if score > 0 {
			break
		}
	}
	assert.Equal(t, 5434, score)
}




