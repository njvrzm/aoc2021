package thirteen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayThirteenPartOne(t *testing.T) {
	testCases := []struct {
		path     string
		expected int
	}{
		{path: "testdata/example", expected: 17},
		{path: "../../inputs/thirteen", expected: 695},
	}
	for _, tt := range testCases {
		paper, folds := GetInput(tt.path)
		paper = paper.Fold(folds[0])
		assert.Equal(t, tt.expected, len(paper))
	}
}
func TestDayThirteenPartTwo(t *testing.T) {
	paper, folds := GetInput("../../inputs/thirteen")
	paper = paper.Apply(folds)
	paper.Read()
}
